package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gen0cide/laforge/ent"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	// s := "blah\nblah💔"
	// fmt.Printf("%s\n---\n", s)
	// s = strings.ReplaceAll(s, "\n", "🔥")
	// fmt.Printf("%s\n---\n", s)
	// s = strings.ReplaceAll(s, "🔥", "\n")
	// fmt.Printf("%s\n---\n", s)
	logrus.SetLevel(logrus.DebugLevel)
	pgHost, ok := os.LookupEnv("PG_HOST")
	client := &ent.Client{}

	if !ok {
		client = ent.PGOpen("postgresql://laforger:laforge@127.0.0.1/laforge")
	} else {
		client = ent.PGOpen(pgHost)
	}

	ctx := context.Background()
	defer ctx.Done()
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// hosts, err := client.Host.Query().All(ctx)
	// if err != nil {
	// 	log.Fatalf("error querying env: %v", err)
	// }

	uuid, _ := uuid.Parse("0f8ce3a7-2d7d-4791-a25b-60d5afbdfdf9")
	build := client.Build.GetX(ctx, uuid)
	teams := build.QueryBuildToTeam().AllX(ctx)

	for _, teamer := range teams {
		ph, err := client.ProvisionedHost.Query().Where(provisionedhost.And(
			provisionedhost.HasProvisionedHostToProvisionedNetworkWith(provisionednetwork.HasProvisionedNetworkToTeamWith(team.IDEQ(teamer.ID))),
			provisionedhost.AddonTypeEQ(provisionedhost.AddonTypeDNS),
		)).All(ctx)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		fmt.Println(ph)
	}

	// build, err := env.QueryEnvironmentToBuild().Order(ent.Desc(build.FieldRevision)).First(ctx)
	// if err != nil {
	// 	log.Fatalf("error w/ build: %v", err)
	// }

	// rootPlan, err := build.QueryBuildToPlan().Where(plan.TypeEQ(plan.TypeStartBuild)).First(ctx)
	// uuid, _ := uuid.Parse("fa4018ac-31f9-4165-a958-d901cc55a96e")
	// rootPlan, err := client.Plan.Query().Where(plan.IDEQ(uuid)).Only(ctx)
	// if err != nil {
	// 	log.Fatalf("error w/ rootPlan: %v", err)
	// }
	// prevPlans, err := rootPlan.PrevPlan(ctx)
	// if err != nil {
	// 	log.Fatalf("error w/ rootPlan: %v", err)
	// }

	// // planPath := ""
	// var wg sync.WaitGroup
	// for _, prevPlan := range prevPlans {
	// 	fmt.Printf("%s\n", prevPlan.ID)
	// 	// wg.Add(1)
	// 	// go Traverse(ctx, planPath, prevPlan, &wg)
	// }
	// wg.Wait()
}

func Traverse(ctx context.Context, planPath string, entPlan *ent.Plan, wg *sync.WaitGroup) {
	defer wg.Done()

	switch entPlan.Type {
	case plan.TypeStartBuild:
		build, err := entPlan.PlanToBuild(ctx)
		if err != nil {
			return
		}
		planPath += fmt.Sprintf("/%s", build.ID)
	case plan.TypeStartTeam:
		team, err := entPlan.PlanToTeam(ctx)
		if err != nil {
			return
		}
		planPath += fmt.Sprintf("/Team%d", team.TeamNumber)
	case plan.TypeProvisionNetwork:
		pnet, err := entPlan.PlanToProvisionedNetwork(ctx)
		if err != nil {
			return
		}
		planPath += fmt.Sprintf("/%s", pnet.Name)
	case plan.TypeProvisionHost:
		phost, err := entPlan.PlanToProvisionedHost(ctx)
		if err != nil {
			return
		}
		planPath += fmt.Sprintf("/%s", phost.SubnetIP)
	case plan.TypeExecuteStep:
		step, err := entPlan.PlanToProvisioningStep(ctx)
		if err != nil {
			return
		}
		planPath += fmt.Sprintf("/%s", step.Type)
	default:
		break
	}
	// fmt.Println(planPath)

	nextPlans, err := entPlan.QueryNextPlan().All(ctx)
	if err != nil {
		return
	}
	if len(nextPlans) == 0 {
		fmt.Println(planPath)
	}

	var nextWg sync.WaitGroup
	for _, nextPlan := range nextPlans {
		nextWg.Add(1)
		go Traverse(ctx, planPath, nextPlan, &nextWg)
	}
	nextWg.Wait()
}
