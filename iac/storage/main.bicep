@description('The azure region/location')
param location string

@metadata({
  suffix: 'The resource suffix'
  suffixCondensed: 'The resource condensed suffix'
})
param naming object

@description('The tags to apply to the resource')
param tags object

resource storage 'Microsoft.Storage/storageAccounts@2021-08-01' = {
  name: 'sa${naming.suffixCondensed}'
  kind: 'StorageV2'
  location: location
  tags: tags
  sku: {
    name: 'Standard_RAGRS'
  }
  properties: {
    accessTier: 'Hot'
    allowBlobPublicAccess: false
    minimumTlsVersion: 'TLS1_2'
    publicNetworkAccess: 'Disabled'
    supportsHttpsTrafficOnly: true
  }
}
