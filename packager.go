// Type of packages
package omg8583

type pkg struct {
	Length int
	Name   string
	Type   string
}

var t = map[int]pkg{
	0: {
		Length: 4,
		Name:   "Message Type Indicator",
		Type:   "n",
	},
	1: {
		Length: 16,
		Name:   "Bitmap",
		Type:   "b",
	},
	2: {
		Length: 19,
		Name:   "Primary Account Number",
		Type:   "lln",
	},
	3: {
		Length: 6,
		Name:   "Processing Code",
		Type:   "n",
	},
	4: {
		Length: 12,
		Name:   "Amount, Transaction",
		Type:   "n",
	},
	5: {
		Length: 12,
		Name:   "Amount, Settlement",
		Type:   "n",
	},
	6: {
		Length: 12,
		Name:   "Amount, Cardholder Billing",
		Type:   "n",
	},
	7: {
		Length: 10,
		Name:   "Transmission Date and Time",
		Type:   "n",
	},
	8: {
		Length: 8,
		Name:   "Amount, Cardholder Billing Fee",
		Type:   "n",
	},
	9: {
		Length: 8,
		Name:   "Conversion Rate, Settlement",
		Type:   "n",
	},
	10: {
		Length: 8,
		Name:   "Conversion Rate, Cardholder Billing",
		Type:   "n",
	},
	11: {
		Length: 6,
		Name:   "System Trace Audit Number",
		Type:   "n",
	},
	12: {
		Length: 6,
		Name:   "Time, Local Transaction",
		Type:   "n",
	},
	13: {
		Length: 4,
		Name:   "Date, Local Transaction",
		Type:   "n",
	},
	15: {
		Length: 4,
		Name:   "Date, Expiration", //Date, Settlement
		Type:   "n",
	},
	32: {
		Length: 11,
		Name:   "Acquiring Institution Ident Code",
		Type:   "lln",
	},
	37: {
		Length: 12,
		Name:   "Retrieval Reference Number",
		Type:   "an",
	},
	38: {
		Length: 6,
		Name:   "Authorization identification response",
		Type:   "an",
	},
	39: {
		Length: 2,
		Name:   "Response code",
		Type:   "an",
	},
	41: {
		Length: 8,
		Name:   "Card Acceptor Terminal Identification",
		Type:   "ans",
	},
	49: {
		Length: 3,
		Name:   "Currency code, transaction",
		Type:   "a",
	},
	54: {
		Length: 120,
		Name:   "Additional amounts",
		Type:   "lllan",
	},
	70: {
		Length: 3,
		Name:   "Network Management Information Code",
		Type:   "n",
	},
}
