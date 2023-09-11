==Long==
# Export
Start an Export job in IdentityNow

Valid Types that can be included or excluded are:
 - ACCESS_PROFILE
 - ACCESS_REQUEST_CONFIG
 - ATTR_SYNC_SOURCE_CONFIG
 - AUTH_ORG
 - CAMPAIGN_FILTER
 - FORM_DEFINITION
 - GOVERNANCE_GROUP
 - IDENTITY_OBJECT_CONFIG
 - IDENTITY_PROFILE
 - LIFECYCLE_STATE
 - NOTIFICATION_TEMPLATE
 - PASSWORD_POLICY
 - PASSWORD_SYNC_GROUP
 - PUBLIC_IDENTITIES_CONFIG
 - ROLE
 - RULE
 - SERVICE_DESK_INTEGRATION
 - SOD_POLICY
 - SOURCE
 - TRANSFORM
 - TRIGGER_SUBSCRIPTION
 - WORKFLOWS
====

==Example==
```bash
sail spconfig export --include WORKFLOWS --include SOURCE
sail spconfig export --include SOURCE --wait
sail spconfig export --include TRANSFORM --objectOptions '{
    "TRANSFORM": {
      "includedIds": [],
      "includedNames": [
        "Remove Diacritical Marks",
        "Common - Location Lookup"
      ]
    }
  }'
```
====