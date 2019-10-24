/*
Package upgrade provides a Cosmos SDK module that can be used for smoothly upgrading a live Cosmos chain to a
new software version. It accomplishes this by providing a BeginBlocker hook that prevents the blockchain state
machine from proceeding once a pre-defined upgrade block time or height has been reached. The module does not prescribe
anything regarding how governance decides to do an upgrade, but just the mechanism for coordinating the upgrade safely.
Without software support for upgrades, upgrading a live chain is risky because all of the validators need to pause
their state machines at exactly the same point in the process. If this is not done correctly, there can be state
inconsistencies which are hard to recover from.

General Workflow

Let's assume we are running v0.34.0 of our software in our testnet and want to upgrade to v0.36.0.
How would this look in practice? First of all, we want to finalize the v0.36.0 release candidate
and there install a specially named upgrade handler (eg. "testnet-v2" or even "v0.36.0"). Once
this code is public, we can have a governance vote to approve this upgrade at some future block time
or block height (e.g. 200000). This is known as an upgrade.Plan. The v0.34.0 code will not know of this
handler, but will continue to run until block 200000, when the plan kicks in at BeginBlock. It will check
for existence of the handler, and finding it missing, know that it is running the obsolete software,
and gracefully exit.

Generally the application binary will restart on exit, but then will execute this BeginBlocker
again and exit, causing a restart loop. Either the operator can manually install the new software,
or you can make use of an external watcher daemon to possibly download and then switch binaries,
also potentially doing a backup. An example of such a daemon is https://github.com/regen-network/cosmosd/
described below under "Automation".

When the binary restarts with the upgraded version (here v0.36.0), it will detect we have registered the
"testnet-v2" upgrade handler in the code, and realize it is the new version. It then will run the upgrade handler
and *migrate the database in-place*. Once finished, it marks the upgrade as done, and continues processing
the rest of the block as normal. Once 2/3 of the voting power has upgraded, the blockchain will immediately
resume the consensus mechanism. If the majority of operators add a custom `do-upgrade` script, this should
be a matter of minutes and not even require them to be awake at that time.

Integrating With An App

Setup an upgrade Keeper for the app and then define a BeginBlocker that calls the upgrade
keeper's BeginBlocker method:
    func (app *myApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
    	app.upgradeKeeper.BeginBlocker(ctx, req)
    	return abci.ResponseBeginBlock{}
    }

The app must then integrate the upgrade keeper with its governance module as appropriate. The governance module
should call ScheduleUpgrade to schedule an upgrade and ClearUpgradePlan to cancel a pending upgrade.

Performing Upgrades

Upgrades can be scheduled at either a predefined block height or time. Once this block height or time is reached, the
existing software will cease to process ABCI messages and a new version with code that handles the upgrade must be deployed.
All upgrades are coordinated by a unique upgrade name that cannot be reused on the same blockchain. In order for the upgrade
module to know that the upgrade has been safely applied, a handler with the name of the upgrade must be installed.
Here is an example handler for an upgrade named "my-fancy-upgrade":
	app.upgradeKeeper.SetUpgradeHandler("my-fancy-upgrade", func(ctx sdk.Context, plan upgrade.Plan) {
		// Perform any migrations of the state store needed for this upgrade
	})

This upgrade handler performs the dual function of alerting the upgrade module that the named upgrade has been applied,
as well as providing the opportunity for the upgraded software to perform any necessary state migrations.

Halt Behavior

Before halting the ABCI state machine in the BeginBlocker method, the upgrade module will log an error
that looks like:
	UPGRADE "<Name>" NEEDED at height <NNNN>: <Info>
where Name are Info are the values of the respective fields on the upgrade Plan.

To perform the actual halt of the blockchain, the upgrade keeper simply panics which prevents the ABCI state machine
from proceeding but doesn't actually exit the process. Exiting the process can cause issues for other nodes that start
to lose connectivity with the exiting nodes, thus this module prefers to just halt but not exit.

Automation and Plan.Info

We have deprecated calling out to scripts, instead with propose https://github.com/regen-network/cosmosd
as a model for a watcher daemon that can launch gaiad as a subprocess and then read the upgrade log message
to swap binaries as needed. You can pass in information into Plan.Info according to the format
specified here https://github.com/regen-network/cosmosd/blob/master/README.md#auto-download .
This will allow a properly configured cosmsod daemon to auto-download new binaries and auto-upgrade.
As noted there, this is intended more for full nodes than validators.

*/
package upgrade
