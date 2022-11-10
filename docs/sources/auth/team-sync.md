+++
title = "Team Sync"
description = "Plutono Team Sync"
keywords = ["plutono", "auth", "documentation"]
aliases = ["/docs/plutono/latest/auth/saml/"]
weight = 1200
+++

# Team sync

With Team Sync, you can set up synchronization between your auth provider's teams and teams in Plutono. This enables LDAP or GitHub OAuth users which are members
of certain teams/groups to automatically be added/removed as members to certain teams in Plutono. Currently the synchronization will only happen every
time a user logs in, unless LDAP is used together with active background synchronization that was added in Plutono 6.3.

{{< figure src="/static/img/docs/enterprise/team_members_ldap.png" class="docs-image--no-shadow docs-image--right" max-width= "600px" >}}

Plutono keeps track of all synchronized users in teams and you can see which users have been synchronized in the team members list, see `LDAP` label in screenshot.
This mechanism allows Plutono to remove an existing synchronized user from a team when its LDAP group membership (for example) changes. This mechanism also enables you to manually add a user as member of a team and it will not be removed when the user signs in. This gives you flexibility to combine LDAP group memberships and Plutono team memberships.

<div class="clearfix"></div>

> Team Sync is only available in Plutono Enterprise.  For more information, refer to [Team sync]({{< relref "../enterprise/team-sync.md" >}}) in [Plutono Enterprise]({{< relref "../enterprise" >}}).
