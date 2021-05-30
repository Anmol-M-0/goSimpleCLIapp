package main

import "github.com/Anmol-M-0/goSimpleCLIapp/app/cliapp/customerCredit/customerCreditApplication"

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	// defer profile.Start(profile.GoroutineProfile).Stop()
	// defer profile.Start(profile.MemProfileAllocs).Stop()
	// defer profile.Start(profile.MemProfileHeap).Stop()
	// defer profile.Start(profile.TraceProfile).Stop()
	// defer profile.Start().Stop()
	customerCreditApplication.RunCustomerCreditApp()
}
