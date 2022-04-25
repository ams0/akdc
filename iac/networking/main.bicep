@description('The azure region/location')
param location string

@description('The tags to apply to the resource')
param tags object

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

resource fleet_dns 'Microsoft.Network/dnsZones@2018-05-01' existing = {
  name: dns.name
  scope: resourceGroup(dns.resourceGroup)
}

resource fleet_domain 'Microsoft.DomainRegistration/domains@2021-03-01' existing = {
  name: domain.name
  scope: resourceGroup(domain.resourceGroup)
}

resource flee_certificate 'Microsoft.CertificateRegistration/certificateOrders@2021-03-01' existing = {
  name: ssl.name
  scope: resourceGroup(ssl.resourceGroup)
}
