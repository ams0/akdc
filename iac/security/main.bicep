@description('The azure region/location')
param location string

@metadata({
  suffix: 'The resource suffix'
  suffixCondensed: 'The resource condensed suffix'
})
param naming object

@description('The tags to apply to the resource')
param tags object

resource keyvault 'Microsoft.KeyVault/vaults@2021-11-01-preview' = {
  name: 'kv${naming.suffixCondensed}'
  location: location
  tags: tags
  properties: {
    tenantId: subscription().tenantId
    accessPolicies: [
      {
        objectId: vmIdentity.properties.principalId
        tenantId: vmIdentity.properties.tenantId
        permissions: {
          secrets: [
            'get'
            'list'
          ]
        }
      }
      {
        objectId: adminIdentity.properties.principalId
        tenantId: adminIdentity.properties.tenantId
        permissions: {
          secrets: [
            'all'
          ]
          certificates: [
            'all'
          ]
          keys: [
            'all'
          ]
          storage: [
            'all'
          ]
        }
      }
    ]
    sku: {
      family: 'A'
      name: 'standard'
    }
  }
}

resource vmIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2018-11-30' = {
  name: 'mi-vm-${naming.suffix}'
  location: location
  tags: tags
}

resource adminIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2018-11-30' = {
  name: 'mi-admin-${naming.suffix}'
  location: location
  tags: tags
}
