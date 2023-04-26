package main

const SampleJSON = `{
	"data": {
		"payment": {
			"cc": "{{cc}}",
			"cvv": "{{cvv}}",
			"exp": "07/25"
		},
		"shopper": {
			"name": "John Smith",
			"email": "john@example.com",
			"phone": "555-123-4567"
		},
		"cart": {
			"amount": 1234,
			"items": [
				{
					"name": "t-shirt",
					"sku": "ABCDE",
					"quantity": 1,
					"amount": 1234
				}
			]
		}
	}
}`

const SampleXML = `{
	<data>
		<payment>
			<cc>{{cc}}</cc>
			<cvv>{{cvv}}</cvv>
			<exp>07/25</exp>
		</payment>
		<shopper>
			<name>John smith</name>
			<email>john@example.com</email>
			<phone>555-123-4567</phone>
		</shopper>
		<cart>
			<amount>1234</amount>
			<items>
				<item>
					<name>t-shirt</name>
					<sku>ABCDE</sku>
					<quantity>1</quantity>
					<amount>1234</amount>
				</item>
			</items>
		</cart>
	</data>
}`
