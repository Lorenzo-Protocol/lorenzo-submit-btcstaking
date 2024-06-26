package btc

/*
{
        "txid": "fbcfe758037013bd2bdab2edc0cb35d843dacc92f56f280a6080b614c5f0202e",
        "version": 2,
        "locktime": 0,
        "vin": [
            {
                "txid": "cf72a95f2d5a0c33a80332ff2de4f7448cb85381b98f487dbad90b197efa67ca",
                "vout": 2,
                "prevout": {
                    "scriptpubkey": "5120099f46cbdfeeeff497b18f16757da503025da16d6bb3eb0108d48810c310bcf4",
                    "scriptpubkey_asm": "OP_PUSHNUM_1 OP_PUSHBYTES_32 099f46cbdfeeeff497b18f16757da503025da16d6bb3eb0108d48810c310bcf4",
                    "scriptpubkey_type": "v1_p2tr",
                    "scriptpubkey_address": "tb1ppx05dj7lamhlf9a33ut82ld9qvp9mgtddwe7kqgg6jyppscshn6qm2926a",
                    "value": 297847
                },
                "scriptsig": "",
                "scriptsig_asm": "",
                "witness": [
                    "184b64ccec114ea2060da7fe5a91f537a775d80c0f7084bab8d3e58f80254852a105c54f727608f5de226e58544c710fe417b0d85ad4f2079886058e11cddf0b"
                ],
                "is_coinbase": false,
                "sequence": 4294967293
            }
        ],
        "vout": [
            {
                "scriptpubkey": "6a2a307830353334416245363263323365364632446332323934433762343645363334303634333334366165",
                "scriptpubkey_asm": "OP_RETURN OP_PUSHBYTES_42 307830353334416245363263323365364632446332323934433762343645363334303634333334366165",
                "scriptpubkey_type": "op_return",
                "value": 0
            },
            {
                "scriptpubkey": "5120da3c374fec13f4c16941c08dd96467772f9070e7d112994ac4ccd9899116963c",
                "scriptpubkey_asm": "OP_PUSHNUM_1 OP_PUSHBYTES_32 da3c374fec13f4c16941c08dd96467772f9070e7d112994ac4ccd9899116963c",
                "scriptpubkey_type": "v1_p2tr",
                "scriptpubkey_address": "tb1pmg7rwnlvz06vz62pczxajer8wuhequ886yffjjkyenvcnygkjc7q7jc2dj",
                "value": 1000
            },
            {
                "scriptpubkey": "5120099f46cbdfeeeff497b18f16757da503025da16d6bb3eb0108d48810c310bcf4",
                "scriptpubkey_asm": "OP_PUSHNUM_1 OP_PUSHBYTES_32 099f46cbdfeeeff497b18f16757da503025da16d6bb3eb0108d48810c310bcf4",
                "scriptpubkey_type": "v1_p2tr",
                "scriptpubkey_address": "tb1ppx05dj7lamhlf9a33ut82ld9qvp9mgtddwe7kqgg6jyppscshn6qm2926a",
                "value": 295847
            }
        ],
        "size": 258,
        "weight": 828,
        "fee": 1000,
        "status": {
            "confirmed": true,
            "block_height": 2584224,
            "block_hash": "0000000000237ebaaa141dd9bd63f9b55c2e96752da19bfb96e490c5b8525164",
            "block_time": 1711708278
        }
    }
*/

type Vout struct {
	ScriptPubKey        string `json:"scriptpubkey"`
	ScriptPubKeyAsm     string `json:"scriptpubkey_asm"`
	ScriptPubKeyType    string `json:"scriptpubkey_type"`
	ScriptPubKeyAddress string `json:"scriptpubkey_address"`
	Value               int    `json:"value"`
}

type Vin struct {
	Txid         string   `json:"txid"`
	Vout         int      `json:"vout"`
	Prevout      Vout     `json:"prevout"`
	ScriptSig    string   `json:"scriptsig"`
	ScriptSigAsm string   `json:"scriptsig_asm"`
	Witness      []string `json:"witness"`
	IsCoinbase   bool     `json:"is_coinbase"`
	Sequence     int      `json:"sequence"`
}

type BtcTx struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
	Size     int    `json:"size"`
	Weight   int    `json:"weight"`
	Fee      int    `json:"fee"`
	Status   struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight int    `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   int    `json:"block_time"`
	} `json:"status"`
}
