@description('Whether or not to register a domain and generate wildcard TLS certificate')
param registerDomain bool = false

@description('The azure region/location')
param location string

@metadata({
  suffix: 'The resource suffix'
  suffixCondensed: 'The resource condensed suffix'
})
param naming object

@description('The contact infomration used for WHOIS domain registration')
param contact object = {
  fistName: 'Wallace'
  lastName: 'Breza'
  email: 'wallace@breza.me'
  phone: '4257041620'
  ipAddress: '67.168.26.44'
}

@description('The tags to apply to the resource')
param tags object

param baseTime string = utcNow()

var domainName = '${naming.suffix}.com'

resource dns 'Microsoft.Network/dnsZones@2018-05-01' = {
  name: domainName
  location: location
  tags: tags
  properties: {
    zoneType: 'Public'
  }
}

resource domain 'Microsoft.DomainRegistration/domains@2021-03-01' = if (registerDomain) {
  name: domainName
  location: location
  tags: tags
  properties: {
    privacy: true
    autoRenew: false
    dnsZoneId: dns.id
    consent: {
      agreedBy: contact.ipAddress
      agreedAt: baseTime
    }
    contactAdmin: {
      email: contact.email
      nameFirst: contact.firstName
      nameLast: contact.lastName
      phone: contact.phone
    }
    contactBilling: {
      email: contact.email
      nameFirst: contact.firstName
      nameLast: contact.lastName
      phone: contact.phone
    }
    contactTech: {
      email: contact.email
      nameFirst: contact.firstName
      nameLast: contact.lastName
      phone: contact.phone
    }
    contactRegistrant: {
      email: contact.email
      nameFirst: contact.firstName
      nameLast: contact.lastName
      phone: contact.phone
    }
  }
}

resource certificate 'Microsoft.CertificateRegistration/certificateOrders@2021-03-01' = if (registerDomain) {
  name: 'crt-${naming.suffix}'
  location: location
  tags: tags
  properties: {
    autoRenew:false
    validityInYears:1
    productType: 'StandardDomainValidatedWildCardSsl'
    distinguishedName: 'CN=*.${domainName}'
  }
}
