# Yet another billing system

## Purpose

Starting with a list of main entieties to be included in a API for a very generic
billing system.

## The ontology

### Event

It is a potential change in the system state. It can be CRUD on any of the entieties
listed below.

Each event has na UUID as its internal key. It can be tagged and sourced ie. parsed
so that its type and contents can be decoded.

Events live on data bus where they are handled by hooks. Handlng means that a specific
API object is requested by CRUD method.

### Label

Virtually anything as it is referencing any object by an UUID.

A tag has a name, a color, may be some other properties.

### Service

It allows an Account holder to generate Usage to be chaerged by a Provider.

A Service has a name, a Tax UUID, an Account UUID.

### Usage

It is an Event of a Usage of a Service by an Account.

A Usage references a Service by UUID.

### Fee

Is is a template for a Charge to be included in an Invoice produced in a Billing run.

A Fee references a Service by UUID.

### Subscription

It grants rights to an Account holder for usage of particular Service.

A Subscription has a Fee template. It is used to produce a Charge in a Billing run.

It has a name. It references a Fee by UUID.

### Provider

It produces a Charge for the holder of an Account for Usage of a Service

It has a name, a address and a tax identification data.

### Account

It is a representation of a physical holder of right to use a Service.

It has a name, an address, etc.

### Billing

It is an Event of a Billing of a Charge for a Service for an Account holder.

### Charge

It is an Event of Billing of a Fee for a Usage of Service on an Account holder.

### Invoice

It is a total sum of a Charges for an Usage of a Service for an Account holder produced by a Provider
in the scope a Billing Event.

It optionally reference another Invoice being its correction by an UUID.

### Ledger

An Charge or a Payment references an Ledger by UUID.

### Tax

It defines final tax to be calculated. Multiple taxes may be applied to an Invoice.

### Payment

A Payment may produce a Cash on account referencing an Invoice by UUID.

### Cash

A Cash references bot an Invoice and a Payment by their UUID.
