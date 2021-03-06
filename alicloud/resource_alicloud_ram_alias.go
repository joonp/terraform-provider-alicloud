package alicloud

import (
	"fmt"

	"github.com/denverdino/aliyungo/ram"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlicloudRamAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudRamAliasCreate,
		Read:   resourceAlicloudRamAliasRead,
		Delete: resourceAlicloudRamAliasDelete,

		Schema: map[string]*schema.Schema{
			"account_alias": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRamAlias,
			},
		},
	}
}

func resourceAlicloudRamAliasCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AliyunClient).ramconn

	args := ram.AccountAliasRequest{
		AccountAlias: d.Get("account_alias").(string),
	}

	if _, err := conn.SetAccountAlias(args); err != nil {
		return fmt.Errorf("SetAccountAlias got an error: %#v", err)
	}

	d.SetId(args.AccountAlias)
	return resourceAlicloudRamAliasRead(d, meta)
}

func resourceAlicloudRamAliasRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AliyunClient).ramconn

	response, err := conn.GetAccountAlias()
	if err != nil {
		return fmt.Errorf("GetAccountAlias got an error: %#v", err)
	}

	d.Set("account_alias", response.AccountAlias)
	return nil
}

func resourceAlicloudRamAliasDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AliyunClient).ramconn

	if _, err := conn.ClearAccountAlias(); err != nil {
		return fmt.Errorf("ClearAccountAlias got an error: %#v", err)
	}
	return nil
}
