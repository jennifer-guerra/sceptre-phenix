package vm

import (
	"fmt"
	"phenix/util/mm/mmcli"

	"github.com/activeshadow/libminimega/miniclient"
)

func checkFirstError(response_channel chan *miniclient.Response) error {
	for element := range response_channel {
		for _, response := range element.Resp {
			if response.Error != "" {
				return fmt.Errorf(response.Error)
			}
		}
	}
	return nil
}

func StartScreenRecord(expNamespace, vmName, fileName string) error {
	if expNamespace == "" {
		return fmt.Errorf("no experiment name provided")
	}

	if vmName == "" {
		return fmt.Errorf("no VM name provided")
	}

	if fileName == "" {
		return fmt.Errorf("no output file provided")
	}

	vm, err := Get(expNamespace, vmName)
	if err != nil {
		return fmt.Errorf("getting VM details: %w", err)
	}

	if !vm.Running {
		return fmt.Errorf("VM is not running")
	}

	// Select active namespace
	cmd := mmcli.Command{
		Command: fmt.Sprint("namespace ", expNamespace),
	}
	if err := checkFirstError(mmcli.Run(&cmd)); err != nil  {
		return err
	}

	// Start VNC recording on VM
	cmd = mmcli.Command{
		Command: fmt.Sprint("vnc record fb ", vmName, fileName),
	}
	if err := checkFirstError(mmcli.Run(&cmd)); err != nil  {
		return err
	}

	return nil
}

func StopScreenRecord(expNamespace, vmName string) error {
	if expNamespace == "" {
		return fmt.Errorf("no experiment name provided")
	}

	if vmName == "" {
		return fmt.Errorf("no VM name provided")
	}
	
	vm, err := Get(expNamespace, vmName)
	if err != nil {
		return fmt.Errorf("getting VM details: %w", err)
	}

	if !vm.Running {
		return fmt.Errorf("VM is not running")
	}

	// Select active namespace
	cmd := mmcli.Command{
		Command: fmt.Sprint("namespace ", expNamespace),
	}
	if err := checkFirstError(mmcli.Run(&cmd)); err != nil  {
		return err
	}

	// Stop VNC recording on VM
	cmd = mmcli.Command{
		Command: fmt.Sprint("vnc stop ", vmName),
	}
	if err := checkFirstError(mmcli.Run(&cmd)); err != nil  {
		return err
	}

	return nil
}

func IsRecording(vmName string) bool {

	return true
}
