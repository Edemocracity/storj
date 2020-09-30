# Planned Downtime 

## Background

With the completion of (TODO link) [storage node downtime tracking with audits](), we will soon be penalizing nodes for downtime through suspension and disqualification. This is good news for durability because up until this point, the only incentive for an SNO to keep their node from going offline is a lower payout - there were no long-term consequences until the new downtime tracking system.

However, there are a variety of situations where an SNO may know in advance that they want to take their node(s) offline for a period of time. For example, if a node operator is moving and runs their node out of their home, they will need to take their node offline in order to move it. If an operator is running nodes out of a datacenter, that datacenter may have planned maintenance scheduled.

"Unplanned downtime" describes any storage node downtime that does not fall under "planned downtime" as described in this document. Unplanned downtime might be caused by a variety of events including, but not limited to, power outages, deliberate or accidental shutoff of a node, or fatal storagenode configuration problems.

## Design

WIP

Basically the goal from a technical perspective is to just not create audit windows for hours where a node has planned downtime. Preferably in a way where we can tell if a node is currently in planned downtime from the `nodes` table with no other tables. If we do not create audit windows for planned downtime hours, node online scores will not be impacted by any downtime during that time.

So if we have a chore that runs every hour (as long as we can make sure it consistently starts/finishes 5-10 minutes before the beginning of the next hour), that chore could check for upcoming planned downtime, then flag all nodes who should have planned downtime in the next hour. Another chore, set to consistently start/finish 5-10 minutes after the beginning of each hour can check for all recently-finished planned downtime and flag all nodes who are no longer in a state of planned downtime.

We also want to track the total amount of planned downtime for each node. So at least two fields need to be added to the `nodes` table: `planned_downtime` boolean, and `total_planned_downtime_hours` int

New table: planned_downtime

fields: node ID, start hour timestamp, duration (in hours)

Open questions:

* How do we limit planned downtime? I see two main options (but there are probably more): 1. Have a specific number of hours per year that a node can use for planned downtime. Once they go over, they simply stop being able to schedule planned downtime. 2. No strict limit on planned downtime. Rather, financially disincentivize nodes from taking planned downtime. No disincentive for first x hours, but after that, planned downtime can reduce payout
* What should the interface for planned downtime look like?
* How far in advance should planned downtime be scheduled? >24 hours? >1 month? No limitation?
* Can/should we do anything with our knowledge of planned downtime? Does a node in planned downtime count against segment health even though we know it will come back online later? Maybe it counts against segment health, but we do not repair off of it in the event that the segment needs to be repaired.




## Rationale

## Implementation

## Wrapup

