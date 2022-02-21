package handler

import (
	"context"
	"github.com/micro/micro/v3/service/store"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"
)

type CookieStealer struct{}

// StealTeamUnifyCookies is a single request handler called via client.Call or the generated client code
func (e *CookieStealer) StealTeamUnifyCookies(
	ctx context.Context,
	_ *cookiestealer.StealTeamUnifyCookiesRequest,
	rsp *cookiestealer.StealTeamUnifyCookiesResponse,
) error {
	log.Info("Received CookieStealer.StealTeamUnifyCookies request")

	read, err := store.Read("cookie")
	if err != nil {
		log.Errorf("Error getting store cookies: %v", err)
	}
	if len(read) > 0 {
		var cookie string
		err = read[0].Decode(&cookie)
		if err != nil {
			return err
		}
		rsp.Cookies = cookie
		return nil
	}

	awsUserVals, err := config.Get("aws.user", config.Secret(true))
	if err != nil {
		return err
	}
	awsUser := awsUserVals.StringMap(nil)
	if awsUser == nil {
		return errors.InternalServerError("CookieStealer.StealTeamUnifyCookies", "Unable to get AWS credentials")
	}

	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(awsUser["key"], awsUser["secret"], ""),
	))
	if err != nil {
		return err
	}

	client := dynamodb.NewFromConfig(cfg)

	tableName, err := getTable()
	if err != nil {
		return err
	}

	scan, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return err
	}

	if scan.Count == 0 {
		// Launch Task
		err = launchTask(ctx, cfg)
		if err != nil {
			return err
		}

		rsp.Cookies = ""
		rsp.Unready = true
		return nil
	}

	var latestCookie string
	var latestTime uint64 = 0
	for _, item := range scan.Items {
		value, err := strconv.ParseUint(item["time"].(*ddbTypes.AttributeValueMemberN).Value, 10, 64)

		if err != nil {
			return err
		}

		if value > latestTime {
			latestTime = value
			latestCookie = item["cookies"].(*ddbTypes.AttributeValueMemberS).Value
		}
	}

	rsp.Cookies = latestCookie
	rsp.Unready = false

	record := store.NewRecord("cookie", latestCookie)
	record.Expiry = 3 * time.Hour

	err = store.Write(record)
	if err != nil {
		return err
	}

	return nil
}

func getTable() (string, error) {
	tableVals, err := config.Get("aws.cookie-table")
	if err != nil {
		return "", err
	}

	table := tableVals.String("")
	if table == "" {
		return "", errors.InternalServerError("CookieStealer.StealTeamUnifyCookies", "Unable to get DynamoDB table")
	}

	return table, nil
}

type TaskConfig struct {
	ClusterId       string
	SecurityGroupId string
	Subnets         []string
	TaskArn         string
}

func LoadTaskConfig() (TaskConfig, error) {
	vals, err := config.Get("aws.ecs")
	if err != nil {
		return TaskConfig{}, err
	}

	values := vals.StringMap(nil)
	if values == nil {
		return TaskConfig{}, errors.InternalServerError("CookieStealer.StealTeamUnifyCookies", "Unable to get Task configuration")
	}

	subnetVals, err := config.Get("aws.ecs.subnets")
	if err != nil {
		return TaskConfig{}, err
	}

	subnetValues := subnetVals.StringSlice([]string{})

	return TaskConfig{
		ClusterId:       values["clusterId"],
		SecurityGroupId: values["sgId"],
		Subnets:         subnetValues,
		TaskArn:         values["taskArn"],
	}, nil
}

func launchTask(ctx context.Context, cfg aws.Config) error {
	client := ecs.NewFromConfig(cfg)

	taskConfig, err := LoadTaskConfig()
	if err != nil {
		return err
	}

	log.Info(taskConfig)

	input := &ecs.RunTaskInput{
		TaskDefinition: aws.String(taskConfig.TaskArn),
		Cluster:        aws.String(taskConfig.ClusterId),
		Count:          aws.Int32(1),
		LaunchType:     ecsTypes.LaunchTypeFargate,
		NetworkConfiguration: &ecsTypes.NetworkConfiguration{
			AwsvpcConfiguration: &ecsTypes.AwsVpcConfiguration{
				AssignPublicIp: ecsTypes.AssignPublicIpEnabled,
				Subnets:        taskConfig.Subnets,
				SecurityGroups: []string{taskConfig.SecurityGroupId},
			},
		},
	}

	_, err = client.RunTask(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
