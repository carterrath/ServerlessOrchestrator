package elasticcontainerservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	r53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go/aws"
)

func RegisterTaskDefinition(client *ecs.Client, backendName string) (*string, error) {
	resp, err := client.RegisterTaskDefinition(context.TODO(), &ecs.RegisterTaskDefinitionInput{
		Family:                  aws.String("microservice-family"),
		Cpu:                     aws.String("256"),
		Memory:                  aws.String("512"),
		NetworkMode:             ecsTypes.NetworkModeAwsvpc,
		RequiresCompatibilities: []ecsTypes.Compatibility{ecsTypes.CompatibilityFargate},
		ExecutionRoleArn:        aws.String("arn:aws:iam::account-id:role/ecsTaskExecutionRole"),
		ContainerDefinitions: []ecsTypes.ContainerDefinition{
			{
				Name:      aws.String(backendName),
				Image:     aws.String("carterrath/serverless-orchestrator:" + backendName),
				Essential: aws.Bool(true),
				PortMappings: []ecsTypes.PortMapping{
					{
						ContainerPort: aws.Int32(3000),
						HostPort:      aws.Int32(3000),
						Protocol:      ecsTypes.TransportProtocolTcp,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.TaskDefinition.TaskDefinitionArn, nil
}

func CreateService(client *ecs.Client, clusterName, serviceName, taskDefinitionArn string) error {
	_, err := client.CreateService(context.TODO(), &ecs.CreateServiceInput{
		Cluster:        aws.String(clusterName),
		ServiceName:    aws.String(serviceName),
		TaskDefinition: aws.String(taskDefinitionArn),
		DesiredCount:   aws.Int32(1),
		LaunchType:     ecsTypes.LaunchTypeFargate,
		NetworkConfiguration: &ecsTypes.NetworkConfiguration{
			AwsvpcConfiguration: &ecsTypes.AwsVpcConfiguration{
				Subnets:        []string{"subnet-xxx"},
				SecurityGroups: []string{"sg-xxx"},
				AssignPublicIp: ecsTypes.AssignPublicIpEnabled,
			},
		},
	})
	return err
}

func CreateDNSRecord(client *route53.Client, domainName, serviceName, dnsName string) error {
	_, err := client.ChangeResourceRecordSets(context.TODO(), &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String("hosted-zone-id"),
		ChangeBatch: &r53Types.ChangeBatch{
			Changes: []r53Types.Change{
				{
					Action: r53Types.ChangeActionUpsert,
					ResourceRecordSet: &r53Types.ResourceRecordSet{
						Name: aws.String(serviceName + "." + domainName),
						Type: r53Types.RRTypeCname,
						ResourceRecords: []r53Types.ResourceRecord{
							{
								Value: aws.String(dnsName),
							},
						},
						TTL: aws.Int64(300),
					},
				},
			},
		},
	})
	return err
}
