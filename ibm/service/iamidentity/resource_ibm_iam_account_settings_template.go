// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func ResourceIBMIamAccountSettingsTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamAccountSettingsTemplateCreate,
		ReadContext:   resourceIBMIamAccountSettingsTemplateRead,
		UpdateContext: resourceIBMIamAccountSettingsTemplateUpdate,
		DeleteContext: resourceIBMIamAccountSettingsTemplateDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ID of the account where the template resides.",
			},
			"name": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"name", "description", "account_settings"},
				Optional:     true,
				Description:  "The name of the trusted profile template. This is visible only in the enterprise account.",
			},
			"description": {
				Type:         schema.TypeString,
				AtLeastOneOf: []string{"name", "description", "account_settings"},
				Optional:     true,
				Description:  "The description of the trusted profile template. Describe the template for enterprise account users.",
			},
			"account_settings": {
				Type:         schema.TypeList,
				AtLeastOneOf: []string{"name", "description", "account_settings"},
				MaxItems:     1,
				Optional:     true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"restrict_create_service_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "NOT_SET",
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_create_platform_apikey": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "NOT_SET",
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_user_list_visibility": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "NOT_RESTRICTED",
							Description: "Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.",
						},
						"restrict_user_domains": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"realm_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The realm that the restrictions apply to.",
									},
									"invitation_email_allow_patterns": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"restrict_invitation": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.",
									},
								},
							},
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
						},
						"mfa": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"session_expiration_in_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.",
						},
						"session_invalidation_in_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.",
						},
						"max_sessions_per_identity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.",
						},
						"system_access_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.",
						},
						"system_refresh_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.",
						},
						"user_mfa": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of users that are exempted from the MFA requirement of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iam_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The iam_id of the user.",
									},
									"mfa": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
									},
								},
							},
						},
						"restrict_user_domains_account_override": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Defines if enterprise defined domain restrictions can be ignored in favour of the restriction defined at the account level.",
						},
					},
				},
			},
			"template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the the template.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the the template resource.",
			},
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Version of the the template.",
			},
			"committed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				Description: "Committed flag determines if the template is ready for assignment.",
			},
			"history": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "History of the Template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the action was triggered.",
						},
						"iam_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM ID of the identity which triggered the action.",
						},
						"iam_id_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account of the identity which triggered the action.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Action of the history entry.",
						},
						"params": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Params of the history entry.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message which summarizes the executed action.",
						},
					},
				},
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity tag for this templateId-version combination.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud resource name.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Created At.",
			},
			"created_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAMid of the creator.",
			},
			"last_modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template last modified at.",
			},
			"last_modified_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAMid of the identity that made the latest modification.",
			},
		},
	}
}

func resourceIBMIamAccountSettingsTemplateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := d.GetOk("template_id"); ok { // if template_id is present then we need to create a new version of this template instead
		return resourceIBMIamAccountSettingsTemplateCreateVersion(context, d, meta)
	}

	createAccountSettingsTemplateOptions := &iamidentityv1.CreateAccountSettingsTemplateOptions{}
	if _, ok := d.GetOk("account_id"); ok {
		createAccountSettingsTemplateOptions.SetAccountID(d.Get("account_id").(string))
	} else {
		userDetails, _ := meta.(conns.ClientSession).BluemixUserDetails()
		accountID := userDetails.UserAccount
		createAccountSettingsTemplateOptions.SetAccountID(accountID)
	}
	if _, ok := d.GetOk("name"); ok {
		createAccountSettingsTemplateOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createAccountSettingsTemplateOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("account_settings"); ok {
		accountSettingsModel, err := ResourceIBMIamAccountSettingsTemplateMapToTemplateAccountSettings(d.Get("account_settings.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "create", "parse-account_settings").GetDiag()
		}
		createAccountSettingsTemplateOptions.SetAccountSettings(accountSettingsModel)
	}

	accountSettingsTemplateResponse, _, err := iamIdentityClient.CreateAccountSettingsTemplateWithContext(context, createAccountSettingsTemplateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAccountSettingsTemplateWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(buildResourceIdFromTemplateVersion(*accountSettingsTemplateResponse.ID, *accountSettingsTemplateResponse.Version))

	if d.Get("committed").(bool) {
		err := resourceIBMAccountSettingsTemplateCommit(context, d, meta)
		if err != nil {
			log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateCommit failed %s", err)
			return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateCommit failed %s", err))
		}
	}

	return resourceIBMIamAccountSettingsTemplateRead(context, d, meta)
}

func resourceIBMIamAccountSettingsTemplateCreateVersion(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createAccountSettingsTemplateVersionOptions := &iamidentityv1.CreateAccountSettingsTemplateVersionOptions{}

	id, _, err := parseTemplateResourceId(d.Get("template_id").(string))
	if err != nil {
		log.Printf("[DEBUG] resourceIBMIamAccountSettingsTemplateCreateVersion failed %s", err)
		return diag.FromErr(fmt.Errorf("resourceIBMIamAccountSettingsTemplateCreateVersion failed %s", err))
	}

	createAccountSettingsTemplateVersionOptions.SetTemplateID(id)

	if _, ok := d.GetOk("account_id"); ok {
		createAccountSettingsTemplateVersionOptions.SetAccountID(d.Get("account_id").(string))
	} else {
		userDetails, _ := meta.(conns.ClientSession).BluemixUserDetails()
		accountID := userDetails.UserAccount
		createAccountSettingsTemplateVersionOptions.SetAccountID(accountID)
	}
	if _, ok := d.GetOk("name"); ok {
		createAccountSettingsTemplateVersionOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createAccountSettingsTemplateVersionOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("account_settings"); ok {
		accountSettingsModel, err := ResourceIBMIamAccountSettingsTemplateMapToTemplateAccountSettings(d.Get("account_settings.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createAccountSettingsTemplateVersionOptions.SetAccountSettings(accountSettingsModel)
	}

	accountSettingsTemplateResponse, response, err := iamIdentityClient.CreateAccountSettingsTemplateVersionWithContext(context, createAccountSettingsTemplateVersionOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateAccountSettingsTemplateVersionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateAccountSettingsTemplateVersionWithContext failed %s\n%s", err, response))
	}

	d.SetId(buildResourceIdFromTemplateVersion(*accountSettingsTemplateResponse.ID, *accountSettingsTemplateResponse.Version))

	if d.Get("committed").(bool) {
		err := resourceIBMAccountSettingsTemplateCommit(context, d, meta)
		if err != nil {
			log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateCommit failed %s", err)
			return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateCommit failed %s", err))
		}
	}

	return resourceIBMIamAccountSettingsTemplateRead(context, d, meta)
}

func resourceIBMIamAccountSettingsTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAccountSettingsTemplateVersionOptions := &iamidentityv1.GetAccountSettingsTemplateVersionOptions{}

	id, version, err := parseTemplateResourceId(d.Id())
	if err != nil {
		log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateRead failed %s", err)
		return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateRead failed %s", err))
	}

	getAccountSettingsTemplateVersionOptions.SetTemplateID(id)
	getAccountSettingsTemplateVersionOptions.SetVersion(version)

	accountSettingsTemplateResponse, response, err := iamIdentityClient.GetAccountSettingsTemplateVersionWithContext(context, getAccountSettingsTemplateVersionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccountSettingsTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(accountSettingsTemplateResponse.Version) {
		if err = d.Set("version", accountSettingsTemplateResponse.Version); err != nil {
			return diag.FromErr(fmt.Errorf("error setting version: %s", err))
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.AccountID) {
		if err = d.Set("account_id", accountSettingsTemplateResponse.AccountID); err != nil {
			err = fmt.Errorf("Error setting account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-account_id").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.Name) {
		if err = d.Set("name", accountSettingsTemplateResponse.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.Description) {
		if err = d.Set("description", accountSettingsTemplateResponse.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.AccountSettings) {
		accountSettingsMap, err := ResourceIBMIamAccountSettingsTemplateTemplateAccountSettingsToMap(accountSettingsTemplateResponse.AccountSettings)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "account_settings-to-map").GetDiag()
		}
		if err = d.Set("account_settings", []map[string]interface{}{accountSettingsMap}); err != nil {
			err = fmt.Errorf("Error setting account_settings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-account_settings").GetDiag()
		}
	}
	if err = d.Set("id", accountSettingsTemplateResponse.ID); err != nil {
		err = fmt.Errorf("Error setting id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-id").GetDiag()
	}
	if err = d.Set("committed", accountSettingsTemplateResponse.Committed); err != nil {
		err = fmt.Errorf("Error setting committed: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-committed").GetDiag()
	}
	var history []map[string]interface{}
	if !core.IsNil(accountSettingsTemplateResponse.History) {
		for _, historyItem := range accountSettingsTemplateResponse.History {
			historyItemMap, err := EnityHistoryRecordToMap(&historyItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "history-to-map").GetDiag()
			}
			history = append(history, historyItemMap)
		}
	}
	if err = d.Set("history", history); err != nil {
		err = fmt.Errorf("Error setting history: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-history").GetDiag()
	}

	if err = d.Set("entity_tag", accountSettingsTemplateResponse.EntityTag); err != nil {
		err = fmt.Errorf("Error setting entity_tag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-entity_tag").GetDiag()
	}
	if err = d.Set("crn", accountSettingsTemplateResponse.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(accountSettingsTemplateResponse.CreatedAt) {
		if err = d.Set("created_at", accountSettingsTemplateResponse.CreatedAt); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.CreatedByID) {
		if err = d.Set("created_by_id", accountSettingsTemplateResponse.CreatedByID); err != nil {
			err = fmt.Errorf("Error setting created_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-created_by_id").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.LastModifiedAt) {
		if err = d.Set("last_modified_at", accountSettingsTemplateResponse.LastModifiedAt); err != nil {
			err = fmt.Errorf("Error setting last_modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-last_modified_at").GetDiag()
		}
	}
	if !core.IsNil(accountSettingsTemplateResponse.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", accountSettingsTemplateResponse.LastModifiedByID); err != nil {
			err = fmt.Errorf("Error setting last_modified_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "read", "set-last_modified_by_id").GetDiag()
		}
	}

	return nil
}

func resourceIBMIamAccountSettingsTemplateUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateAccountSettingsTemplateVersionOptions := &iamidentityv1.UpdateAccountSettingsTemplateVersionOptions{}

	id, version, err := parseTemplateResourceId(d.Id())
	if err != nil {
		log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateUpdate failed %s", err)
		return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateUpdate failed %s", err))
	}

	updateAccountSettingsTemplateVersionOptions.SetTemplateID(id)
	updateAccountSettingsTemplateVersionOptions.SetVersion(version)
	updateAccountSettingsTemplateVersionOptions.SetIfMatch(d.Get("entity_tag").(string))

	hasChange := false

	if _, ok := d.GetOk("account_id"); ok {
		updateAccountSettingsTemplateVersionOptions.SetAccountID(d.Get("account_id").(string))
	}
	if d.HasChange("name") {
		updateAccountSettingsTemplateVersionOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateAccountSettingsTemplateVersionOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("account_settings") {
		accountSettings, err := ResourceIBMIamAccountSettingsTemplateMapToTemplateAccountSettings(d.Get("account_settings.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "update", "parse-account_settings").GetDiag()
		}
		updateAccountSettingsTemplateVersionOptions.SetAccountSettings(accountSettings)
		hasChange = true
	}

	if hasChange {
		_, _, err = iamIdentityClient.UpdateAccountSettingsTemplateVersionWithContext(context, updateAccountSettingsTemplateVersionOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAccountSettingsTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	if d.HasChange("committed") {
		if d.Get("committed").(bool) {
			err := resourceIBMAccountSettingsTemplateCommit(context, d, meta)
			if err != nil {
				log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateCommit failed %s", err)
				return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateCommit failed %s", err))
			}
		} else {
			return diag.FromErr(fmt.Errorf("A committed template cannot be uncommitted"))
		}
	}

	return resourceIBMIamAccountSettingsTemplateRead(context, d, meta)
}

func resourceIBMIamAccountSettingsTemplateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteAccountSettingsTemplateVersionOptions := &iamidentityv1.DeleteAccountSettingsTemplateVersionOptions{}

	id, version, err := parseTemplateResourceId(d.Id())
	if err != nil {
		log.Printf("[DEBUG] resourceIBMIamAccountSettingsTemplateDelete failed %s", err)
		return diag.FromErr(fmt.Errorf("resourceIBMIamAccountSettingsTemplateDelete failed %s", err))
	}

	deleteAccountSettingsTemplateVersionOptions.SetTemplateID(id)
	deleteAccountSettingsTemplateVersionOptions.SetVersion(version)

	_, err = iamIdentityClient.DeleteAccountSettingsTemplateVersionWithContext(context, deleteAccountSettingsTemplateVersionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAccountSettingsTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func resourceIBMAccountSettingsTemplateCommit(context context.Context, d *schema.ResourceData, meta interface{}) error {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}

	id, version, err := parseTemplateResourceId(d.Id())
	if err != nil {
		return err
	}

	commitAccountSettingsTemplateVersionOptions := iamIdentityClient.NewCommitAccountSettingsTemplateOptions(id, version)
	_, err = iamIdentityClient.CommitAccountSettingsTemplateWithContext(context, commitAccountSettingsTemplateVersionOptions)
	if err != nil {
		return err
	}

	return nil
}

func ResourceIBMIamAccountSettingsTemplateMapToTemplateAccountSettings(modelMap map[string]interface{}) (*iamidentityv1.TemplateAccountSettings, error) {
	model := &iamidentityv1.TemplateAccountSettings{}
	if modelMap["restrict_create_service_id"] != nil && modelMap["restrict_create_service_id"].(string) != "" {
		model.RestrictCreateServiceID = core.StringPtr(modelMap["restrict_create_service_id"].(string))
	}
	if modelMap["restrict_create_platform_apikey"] != nil && modelMap["restrict_create_platform_apikey"].(string) != "" {
		model.RestrictCreatePlatformApikey = core.StringPtr(modelMap["restrict_create_platform_apikey"].(string))
	}
	if modelMap["restrict_user_list_visibility"] != nil && modelMap["restrict_user_list_visibility"].(string) != "" {
		model.RestrictUserListVisibility = core.StringPtr(modelMap["restrict_user_list_visibility"].(string))
	}
	if modelMap["restrict_user_domains"] != nil {
		restrictUserDomains := []iamidentityv1.AccountSettingsUserDomainRestriction{}
		for _, restrictUserDomainsItem := range modelMap["restrict_user_domains"].([]interface{}) {
			restrictUserDomainsItemModel, err := ResourceIBMIamAccountSettingsTemplateMapToAccountSettingsUserDomainRestriction(restrictUserDomainsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			restrictUserDomains = append(restrictUserDomains, *restrictUserDomainsItemModel)
		}
		model.RestrictUserDomains = restrictUserDomains
	}
	if modelMap["allowed_ip_addresses"] != nil && modelMap["allowed_ip_addresses"].(string) != "" {
		model.AllowedIPAddresses = core.StringPtr(modelMap["allowed_ip_addresses"].(string))
	}
	if modelMap["mfa"] != nil && modelMap["mfa"].(string) != "" {
		model.Mfa = core.StringPtr(modelMap["mfa"].(string))
	}
	if modelMap["session_expiration_in_seconds"] != nil && modelMap["session_expiration_in_seconds"].(string) != "" {
		model.SessionExpirationInSeconds = core.StringPtr(modelMap["session_expiration_in_seconds"].(string))
	}
	if modelMap["session_invalidation_in_seconds"] != nil && modelMap["session_invalidation_in_seconds"].(string) != "" {
		model.SessionInvalidationInSeconds = core.StringPtr(modelMap["session_invalidation_in_seconds"].(string))
	}
	if modelMap["max_sessions_per_identity"] != nil && modelMap["max_sessions_per_identity"].(string) != "" {
		model.MaxSessionsPerIdentity = core.StringPtr(modelMap["max_sessions_per_identity"].(string))
	}
	if modelMap["system_access_token_expiration_in_seconds"] != nil && modelMap["system_access_token_expiration_in_seconds"].(string) != "" {
		model.SystemAccessTokenExpirationInSeconds = core.StringPtr(modelMap["system_access_token_expiration_in_seconds"].(string))
	}
	if modelMap["system_refresh_token_expiration_in_seconds"] != nil && modelMap["system_refresh_token_expiration_in_seconds"].(string) != "" {
		model.SystemRefreshTokenExpirationInSeconds = core.StringPtr(modelMap["system_refresh_token_expiration_in_seconds"].(string))
	}
	if modelMap["user_mfa"] != nil {
		userMfa := []iamidentityv1.UserMfa{}
		for _, userMfaItem := range modelMap["user_mfa"].([]interface{}) {
			userMfaItemModel, err := ResourceIBMIamAccountSettingsTemplateMapToUserMfa(userMfaItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			userMfa = append(userMfa, *userMfaItemModel)
		}
		model.UserMfa = userMfa
	}
	if modelMap["restrict_user_domains_account_override"] != nil {
		model.RestrictUserDomainsAccountOverride = core.BoolPtr(modelMap["restrict_user_domains_account_override"].(bool))
	}
	return model, nil
}

func ResourceIBMIamAccountSettingsTemplateMapToAccountSettingsUserDomainRestriction(modelMap map[string]interface{}) (*iamidentityv1.AccountSettingsUserDomainRestriction, error) {
	model := &iamidentityv1.AccountSettingsUserDomainRestriction{}
	model.RealmID = core.StringPtr(modelMap["realm_id"].(string))
	if modelMap["invitation_email_allow_patterns"] != nil {
		invitationEmailAllowPatterns := []string{}
		for _, invitationEmailAllowPatternsItem := range modelMap["invitation_email_allow_patterns"].([]interface{}) {
			invitationEmailAllowPatterns = append(invitationEmailAllowPatterns, invitationEmailAllowPatternsItem.(string))
		}
		model.InvitationEmailAllowPatterns = invitationEmailAllowPatterns
	}
	if modelMap["restrict_invitation"] != nil {
		model.RestrictInvitation = core.BoolPtr(modelMap["restrict_invitation"].(bool))
	}
	return model, nil
}

func ResourceIBMIamAccountSettingsTemplateMapToUserMfa(modelMap map[string]interface{}) (*iamidentityv1.UserMfa, error) {
	model := &iamidentityv1.UserMfa{}
	if modelMap["iam_id"] != nil && modelMap["iam_id"].(string) != "" {
		model.IamID = core.StringPtr(modelMap["iam_id"].(string))
	}
	if modelMap["mfa"] != nil && modelMap["mfa"].(string) != "" {
		model.Mfa = core.StringPtr(modelMap["mfa"].(string))
	}
	return model, nil
}

func ResourceIBMIamAccountSettingsTemplateTemplateAccountSettingsToMap(model *iamidentityv1.TemplateAccountSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestrictCreateServiceID != nil {
		modelMap["restrict_create_service_id"] = *model.RestrictCreateServiceID
	}
	if model.RestrictCreatePlatformApikey != nil {
		modelMap["restrict_create_platform_apikey"] = *model.RestrictCreatePlatformApikey
	}
	if model.RestrictUserListVisibility != nil {
		modelMap["restrict_user_list_visibility"] = *model.RestrictUserListVisibility
	}
	if model.RestrictUserDomains != nil {
		restrictUserDomains := []map[string]interface{}{}
		for _, restrictUserDomainsItem := range model.RestrictUserDomains {
			restrictUserDomainsItemMap, err := AccountSettingsUserDomainRestrictionToMap(&restrictUserDomainsItem)
			if err != nil {
				return modelMap, err
			}
			restrictUserDomains = append(restrictUserDomains, restrictUserDomainsItemMap)
		}
		modelMap["restrict_user_domains"] = restrictUserDomains
	}
	if model.AllowedIPAddresses != nil {
		modelMap["allowed_ip_addresses"] = *model.AllowedIPAddresses
	}
	if model.Mfa != nil {
		modelMap["mfa"] = *model.Mfa
	}
	if model.SessionExpirationInSeconds != nil {
		modelMap["session_expiration_in_seconds"] = *model.SessionExpirationInSeconds
	}
	if model.SessionInvalidationInSeconds != nil {
		modelMap["session_invalidation_in_seconds"] = *model.SessionInvalidationInSeconds
	}
	if model.MaxSessionsPerIdentity != nil {
		modelMap["max_sessions_per_identity"] = *model.MaxSessionsPerIdentity
	}
	if model.SystemAccessTokenExpirationInSeconds != nil {
		modelMap["system_access_token_expiration_in_seconds"] = *model.SystemAccessTokenExpirationInSeconds
	}
	if model.SystemRefreshTokenExpirationInSeconds != nil {
		modelMap["system_refresh_token_expiration_in_seconds"] = *model.SystemRefreshTokenExpirationInSeconds
	}
	if model.UserMfa != nil {
		userMfa := []map[string]interface{}{}
		for _, userMfaItem := range model.UserMfa {
			userMfaItemMap, err := UserMfaResponseToMap(&userMfaItem)
			if err != nil {
				return modelMap, err
			}
			userMfa = append(userMfa, userMfaItemMap)
		}
		modelMap["user_mfa"] = userMfa
	}
	if model.RestrictUserDomainsAccountOverride != nil {
		modelMap["restrict_user_domains_account_override"] = *model.RestrictUserDomainsAccountOverride
	}
	return modelMap, nil
}
