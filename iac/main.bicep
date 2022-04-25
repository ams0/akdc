targetScope = 'subscription'

@description('The organization business unit prefix')
param businessUnit string

@description('The application name')
param appName string

@description('The environment name')
param environment string

@description('The azure region or location')
param location string

@description('The app service domain info')
@metadata({
  name: 'The app service domain name'
  resourceGroup: 'The resource group of the app service domain'
})
param domain object

@description('The app service certificate info')
@metadata({
  name: 'The app service certificate name'
  resourceGroup: 'The resource group of the app service certificate'
})
param ssl object

@description('The DNS zone info')
@metadata({
  name: 'The DNS zone name'
  resourceGroup: 'The resource group of the DNS zone'
})
param dns object

param baseTime string = utcNow()

resource resourceGroup 'Microsoft.Resources/resourceGroups@2020-06-01' = {
  name: 'rg-${resourceSuffix}'
  location: location
}

var naming = {
  suffix: resourceSuffix
  suffixCondensed: resourceSuffixCondensed
}

module networking './networking/main.bicep' = {
  name: 'networking-${baseTime}'
  scope: resourceGroup
  params: {
    location: 'global'
    domain: domain
    ssl: ssl
    dns: dns
    tags: tags
  }
}

module security './security/main.bicep' = {
  name: 'security-${baseTime}'
  scope: resourceGroup
  params: {
    location: location
    naming: naming
    tags: tags
  }
}

module storage './storage/main.bicep' = {
  name: 'storage-${baseTime}'
  scope: resourceGroup
  params: {
    location: location
    naming: naming
    tags: tags
  }
}

var locationMap = {
  global: 'gbl'
  eastus: 'eus'
  eastus2: 'eus2'
  southcentralus: 'scus'
  westus2: 'wus2'
  westus3: 'wus3'
  australiaeast: 'aue'
  southeastasia: 'sea'
  northeurope: 'ne'
  swedencentral: 'sec'
  uksouth: 'uks'
  westeurope: 'we'
  centralus: 'cus'
  northcentralus: 'ncus'
  westus: 'wus'
  southafricanorth: 'zan'
  centralindia: 'inc'
  eastasia: 'ea'
  japaneast: 'jpe'
  koreacentral: 'krc'
  canadacentral: 'cac'
  francecentral: 'frc'
  germanywestcentral: 'dewc'
  norwayeast: 'noe'
  switzerlandnorth: 'chn'
  uaenorth: 'aen'
  brazilsouth: 'brs'
  asiapacific: 'asp'
  westcentralus: 'wcus'
  southafricawest: 'zaw'
  australiacentral: 'auc'
  australiacentral2: 'aus2'
  australiasoutheast: 'ause'
  japanwest: 'jpw'
  koreasouth: 'krs'
  southindia: 'ins'
  westindia: 'inw'
  canadaeast: 'cae'
  francesouth: 'frs'
  germanynorth: 'den'
  norwaywest: 'now'
  swedensouth: 'ses'
  switzerlandwest: 'chw'
  ukwest: 'ukw'
  uaecentral: 'aec'
  brazilsoutheast: 'brse'
}

var regionShort = locationMap[location]

var resourceSuffix = '${businessUnit}-${appName}-${environment}-${regionShort}'
var resourceSuffixCondensed = '${businessUnit}${appName}${environment}${regionShort}'
var tags = {
  business_unit: businessUnit
  app_name: appName
  environment: environment
  location: location
}
