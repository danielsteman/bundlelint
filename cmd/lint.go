package cmd

import "fmt"

func ValidateConfigs(bundleConfig *BundleConfig, lintConfig *LintConfig) bool {
	if lintConfig.NotificationsInProd {
		prodTarget, exists := bundleConfig.Targets["prod"]
		if !exists {
			fmt.Println("Validation failed: 'prod' target does not exist.")
			return false
		}

		for jobName, job := range prodTarget.Jobs {
			if job.WebhookNotifications == nil || len(job.WebhookNotifications.OnFailure) == 0 {
				fmt.Printf("Validation failed: Job %q in prod target is missing OnFailure webhook notifications.\n", jobName)
				return false
			}

			valid := false
			for _, webhook := range job.WebhookNotifications.OnFailure {
				if webhook.ID != "" {
					valid = true
					break
				}
			}

			if !valid {
				fmt.Printf("Validation failed: Job %q in prod target has no valid OnFailure webhook IDs.\n", jobName)
				return false
			}
		}
	}
	return true
}

