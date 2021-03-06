package quorum

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Use this resource to construct `extradata` field used in the genesis file.
//
// `istanbul_address` can be referenced from `quorum_bootstrap_node_key` data source or newly created from `quorum_bootstrap_node_key` resources.
func resourceBootstrapIstanbulExtradata() *schema.Resource {
	return &schema.Resource{
		Create: resourceBootstrapIstanbulExtradataCreate,
		Read:   resourceBootstrapIstanbulExtradataRead,
		Delete: resourceBootstrapIstanbulExtradataDelete,

		Schema: map[string]*schema.Schema{
			"istanbul_addresses": {
				Type:        schema.TypeList,
				Description: "list of Istanbul address to construct extradata",
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Required:    true,
				ForceNew:    true,
			},
			"vanity": {
				Type:        schema.TypeString,
				Description: "Vanity Hex Value to be included in the extradata",
				Optional:    true,
				ForceNew:    true,
				Default:     "0x00",
			},
			"extradata": {
				Type:        schema.TypeString,
				Description: "Computed value which can be used in genesis file",
				Computed:    true,
			},
		},
	}
}

func resourceBootstrapIstanbulExtradataCreate(d *schema.ResourceData, _ interface{}) error {
	addresses := d.Get("istanbul_addresses").([]interface{})
	validators := make([]common.Address, len(addresses))
	for idx, rawAddress := range addresses {
		if addr, ok := rawAddress.(string); !ok {
			return fmt.Errorf("expect string element in istanbul_addresses")
		} else {
			validators[idx] = common.HexToAddress(addr)
		}
	}

	ist := &types.IstanbulExtra{
		Validators:    validators,
		Seal:          make([]byte, types.IstanbulExtraSeal),
		CommittedSeal: [][]byte{},
	}
	payload, err := rlp.EncodeToBytes(&ist)
	if err != nil {
		return err
	}

	vanity := d.Get("vanity").(string)
	newVanity, err := hexutil.Decode(vanity)
	if err != nil {
		return err
	}
	if len(newVanity) < types.IstanbulExtraVanity {
		newVanity = append(newVanity, bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity-len(newVanity))...)
	}
	newVanity = newVanity[:types.IstanbulExtraVanity]
	extradata := "0x" + common.Bytes2Hex(append(newVanity, payload...))
	d.Set("extradata", extradata)
	d.SetId(fmt.Sprintf("%d", time.Now().Unix()))
	return nil
}

func resourceBootstrapIstanbulExtradataRead(d *schema.ResourceData, _ interface{}) error {
	return nil
}

func resourceBootstrapIstanbulExtradataDelete(d *schema.ResourceData, _ interface{}) error {
	return nil
}
