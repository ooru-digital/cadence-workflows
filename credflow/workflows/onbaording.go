package workflows

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
)

const (
	// MaxSignalCount      = 4
	DirectorChannel     = "director"
	OrganisationChannel = "organisation"
	WarehouseChannel    = "warehouse"
	BankChannel         = "bank"
)

func OnboardingWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 60,
		StartToCloseTimeout:    time.Minute * 60,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	fmt.Println("**********************************************")
	logger.Info("SignalGreeterMultiLanguageWorkflow started, ")
	fmt.Println("**********************************************")

	signalCount := 0

	//var signalVal string
	var director string
	var organisation string
	var warehouse string
	var bank string

	// signalChan := workflow.GetSignalChannel(ctx, signalName)
	// signalNames are defined in the constants.
	directorChan := workflow.GetSignalChannel(ctx, DirectorChannel)
	orgChan := workflow.GetSignalChannel(ctx, OrganisationChannel)
	warehouseChan := workflow.GetSignalChannel(ctx, WarehouseChannel)
	bankChan := workflow.GetSignalChannel(ctx, BankChannel)

	s := workflow.NewSelector(ctx)
	// -------------------------------Signal Onboarding workflow Director Details step ------------------------------------

	s.AddReceive(directorChan, func(ch workflow.Channel, ok bool) {
		if ok {
			ch.Receive(ctx, &director)
			signalCount += 1
			fmt.Println("signal count in ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(signalCount)
		}
	})
	s.Select(ctx)

	var director_data string
	error_director := workflow.ExecuteActivity(ctx, Directors_Details, director).Get(ctx, &director_data)
	if error_director != nil {
		return error_director
	}
	fmt.Println("*******************Driector details are*******************************")
	fmt.Println(director_data)

	// -------------------------------Signal Onboarding workflow Organisation Details step ------------------------------------

	s.AddReceive(orgChan, func(ch workflow.Channel, ok bool) {
		if ok {
			ch.Receive(ctx, &organisation)
			signalCount += 1
			fmt.Println("signal count in ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(signalCount)
		}
	})
	var data1 string
	error_org := workflow.ExecuteActivity(ctx, Organisation_Details, organisation).Get(ctx, &data1)
	if error_org != nil {
		return error_org
	}
	fmt.Println("")

	// -------------------------------Signal Onboarding workflow Warehouse Details step ------------------------------------

	s.AddReceive(warehouseChan, func(ch workflow.Channel, ok bool) {
		if ok {
			ch.Receive(ctx, &warehouse)
			signalCount += 1
			fmt.Println("signal count in ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(signalCount)
		}
	})
	var data2 string
	error_warehouse := workflow.ExecuteActivity(ctx, Warehouse_Details, warehouse, data2).Get(ctx, &data2)
	if error_warehouse != nil {
		return error_warehouse
	}
	// -------------------------------Signal Onboarding workflow Bank Details step ------------------------------------

	s.AddReceive(bankChan, func(ch workflow.Channel, ok bool) {
		if ok {
			ch.Receive(ctx, &bank)
			signalCount += 1
			fmt.Println("signal count in ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(signalCount)
		}
	})
	var data3 string
	error_bank := workflow.ExecuteActivity(ctx, Bank_Details, bank, data3).Get(ctx, &data3)
	if error_bank != nil {
		return error_bank
	}
	return nil
}

func Directors_Details(ctx context.Context, director, greeting string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Director_details started")
	fmt.Println("**********************************************")
	fmt.Println("**********************************************")
	fmt.Println("Director data is")
	fmt.Println(director)
	// director_name =
	// var form_data string
	return director, nil
}
func Organisation_Details(ctx context.Context, language, greeting string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Organisation_details started")
	fmt.Println("**********************************************")
	fmt.Println("**********************************************")
	var form_data string
	return form_data, nil
}
func Warehouse_Details(ctx context.Context, language, greeting string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Warehouse_details started")
	fmt.Println("**********************************************")
	fmt.Println("**********************************************")
	var form_data string
	return form_data, nil
}
func Bank_Details(ctx context.Context, language, greeting string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Bank_details started")
	fmt.Println("**********************************************")
	fmt.Println("**********************************************")
	var form_data string
	return form_data, nil
}
