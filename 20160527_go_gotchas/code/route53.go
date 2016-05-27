package main

import "os"

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/barnybug/cli53"
)

func main() {
	r53 := route53.New(session.New())
	output, err := r53.ListHostedZones(&route53.ListHostedZonesInput{})
	if err == nil {
		for _, hz := range output.HostedZones {
			cli53.ExportBindToWriter(r53, hz, true, os.Stdout)
		}
	}
}
