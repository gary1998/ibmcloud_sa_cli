# IBM CLOUD SECURITY ADVISOR
Security Advisor CLI

## Usage

The current supported commands are listed below.

```sh
ibmcloud sa - Interact with IBM Cloud Security Advisor.

USAGE:
ibmcloud sa command [arguments...] [command options]

COMMANDS:
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------

create-channel    Create a new notification channel using IBM Cloud Security Advisor Notifications API.
create-note       Create a new note using IBM Cloud Security Advisor Findings API.
delete-channel    Delete a specific notification channel using IBM Cloud Security Advisor Notifications API.
delete-channels   Delete multiple notification channels (bulk) using IBM Cloud Security Advisor Notifications API.
delete-note       Delete a note using IBM Cloud Security Advisor Findings API.
get-channel       Return the details of a specified notification channel using IBM Cloud Security Advisor Notifications API.
get-channels      Return the list of all notification channels using IBM Cloud Security Advisor Notifications API.
get-graph         Get notes and occurrences using GraphQL queries.
get-key           Returns public key using IBM Cloud Security Advisor Notifications API.
get-note          Return the details of a specified note using IBM Cloud Security Advisor Findings API.
get-notes         Return a list of all notes using IBM Cloud Security Advisor Findings API.
get-providers     Returns the list of all providers using IBM Cloud Findings API.
test-channel      Test a specific notification channel using IBM Cloud Security Advisor Notifications API.
update-channel    Update the details of specific notification channel using IBM Cloud Security Advisor Notifications API.
update-note       Updates the details of a specified note using IBM Cloud Security Advisor Findings API.
help, h           Show help

Enter 'ibmcloud sa help [command]' for more information about a command.
```

## Contributing/local setup
- Clone the code

To build run the following
```go
go build
```

The above command will build the binaries.

To install the built plugin locally run the following.
```bash
ibmcloud plugin uninstall security-advisor
ibmcloud plugin install ibmcloud_sa_cli
```

The above assumes you are running on Mac, if you aren't check out the directory for the correct plugin filename
pbcopy <~/.ssh/id_rsa.pub