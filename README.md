# firefly-iii-transaction-relay
I created this Go API Endpoint to accept transaction SMS messages from my bank, parse and extract the information in them to form a Firefly III transaction JSON payload and post it to the API and avoid manually inputting transactions into Firefly.

I say SMS messages, but it can accept literally any form of text message. you can customize what "key phrases" are used to extract the details of the transaction via environment variables, `.env.example` provides an empty template of all the necessary environment variables.

## Installation
You can either run this via Docker, or as an executable. **Make sure to configure environment variables for both methods.**

### Environment Variables
The envornment variables are mandatory for the Program's functions.

```md
- `LISTEN_ADDRESS`: <ADDRESS>:<PORT>, it's almost always best to set it to `0.0.0.0:<PORT>`
- `FIREFLY_ADDRESS`: Adress of your Firefly III Instance.
- `FIREFLY_ACCOUNT`: Name of your Firefly III account to register transactions to.
- `FIREFLY_BUDGET`: Name of your Firefly III budget to register transactions to.
- `FIREFLY_PAT`: generated via `Options>Profile>OAuth`, needed to authenticate to Firefly's API.
- `AUTH_KEY`: Used in combination with `VERIFICATION_KEY` to authenticate incoming requests.
- `VERIFICATION_KEY`: Used in combination with `AUTH_KEY` to authenticate incoming requests.

The next variables are key phrases used to generate the RegEx queries. All are lists seperated by commas

- `BLACKLIST_KEYS`: Used to skip messages containing key phrases, say to ignore OTP messages or subscription transactions, i.e. 'Netflix' or 'OTP'.
- `WITHDRAWAL_KEYS`: Key phrases that indicate a withdrawal transaction, i.e. 'deducted' or 'withdrawn.'
- `DEPOSIT_KEYS`: Key phrases that indicate a deposit transaction, i.e. 'deposited.'
- `CURRENCY_MAP`: A comma seperated list, each element should be in the form of '<KEY PHRASE>:<CODE>', i.e. 'Dollar:USD'
- `AMOUNT_KEYS`: Key phrases that indicate the transcation amount afterwards. Detects a number with optional decimals afterr the key phrase.
- `DATE_KEYS`: Key phrases that indicate a date afterwards, i.e. 'On Date'. This is used in conjunction with `DATE_FORMATS` to grab the date afer the key phrase and validate it with the format. 
- `DATE_FORMATS`: List of Go Compatible Date formats that can appear in your sms messages. this is used to validate dates captured with `DATE_KEYS`, i.e. '2006/01/02 03:04 PM'.
- `WITHDRAWAL_VENDOR_KEYS`: Key phrases to indicate the transaction's other party afterwards. This pattern will capture everything until it detects one of the other keywords, i.e 'From Vendor', 'To Account'. This pattern is used when the transaction's type is detemined to be a withdrawal.
- `DEPOSIT_VENDOR_KEYS`: Key phrases to indicate the transaction's other party afterwards. This pattern will capture everything until it detects one of the other keywords, i.e 'From Vendor', 'To Account'. This pattern is used when the transaction's type is detemined to be a deposit.
```

For clarification on Go compatible date formats, see [this link](https://www.geeksforgeeks.org/go-language/time-formatting-in-golang/).

### Docker
```docker run -p 8080:8080 judexgrim/firefly-iii-transaction-relay:latest --env-file ./.env ```

### Executable
Download the executable suitable for your machine's architecture from the latest release, and run it.

# Usage
- First, you have a `/health` endpoint you can send a `GET` requestto, whcih will return a `200` when the relay is ready.

- To use the relay, you will first need to create a session variable by sending a `POST` request to the `/auth` endpoint, containing a `JSON` body that looks like this:

```json
{
  "auth_key": "<KEY>"
}
```
This will return a token, which needs to be used as an `Authorization` Header in this form: `bearer <TOKEN>`.

- afterwards, you can send a `POST` request to, with the `Authorization` Header and a `JSON` body that looks like this:

```json
{
  "verification_key": "<KEY>",
  "message": "<MESSAGE>"
}
```

- Given you configure your key environments variables appropriately, the relay will then parse the message, extract the needed details, form a valid transaction `JSON` body to send to the Firefly III API.