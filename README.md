# DEBILL

Also known as Durable Event Billing Interface Lazy Layer.

The design of it is impacted by BSCS, BRM, Nobill, Telness, MACH Roaming, etc.

## Purpose

As usual, a PoC to note main entieties and infra components.

## The ontology

Starting with a list of main entieties to be included in a API for a very generic
billing system. It is purely coceptual model.

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

It allows an Account holder to generate Usage to be charged by a Provider if subscribed.

It is assigned as a Subscription with a Fee on a Product.

A Service has a name, a Tax UUID, UoM (Unit of Measure).

A service is available in an Offert. It is provided with a Price. A Fee is created
when a Service is subscribed ie. when a Subscription is done. The Fee gets initailly
a Price from a Offert but it can be modified.

### Usage

It is an Event of a Usage of a Service by an Account.

The Usage Records are stored in the Usage DB. They are DB representations of JSON
structure which itself is a a internal record representation.

A Usage references a Service by UUID.

It has a context like time stamp, measure (in units of a Service), location, destination, account UUID.

The Usage DB is partitioned by account UUID.

### Fee

Is is a template for a Charge to be included in an Invoice produced in a Billing run.

It defines a Price for a Product. It is copied from a Price of the Offfert.

A Fee references a Service, a Product, a Subscription by UUID.

It has a periodicity and its type (one time, weekly, montly, etc.)

### Product

A Product is a Service for which a Subscription may be done. It is offered to an Account
by a Provider.

### Catalog

It is a set of Products. A Catalog contains no Prices, just Product list.

### Offert

An Offert contains Products with a default Price. It has a date since when it is valid.
It has a off date since when it is not valid.

### Subscription

It grants rights to an Account holder for usage of particular Service.

A Subscription has a Fee. It is used to produce a Charge in a Billing run.

It has a name. It references a Fee by UUID.

It has a start data and end date.

It has a status like: Created, Active, Suspended, Deactivated.

A label may be associated with a Subscription and it will be added to the invoice charge item.

### Subscription_History

It stores the event timestamps for a Subscription state changes.

### Provider

It produces a Charge for the holder of an Account for Usage of a Service

It has a name, an Address and a Tax identification data.

### Account

An account has a Subscription for a Service for which a Fee can be assigned.

It has a name, Tags, etc.

### Customer

It is a representation of a physical holder of an Account. It may have many Accounts.

It has a name, an Address, TaxID, etc.

An Invoice is produced for a Customer with Accounts having Subscriptons on Services.

### Address

An Address is a descriptor for a physical location.

### Billing

It is an Event of a Billing of a Charge for a Service for an Account holder - a Customer.

### Charge

It is an Event of Billing of a Fee for a Usage of Service on an Account holder.

### Invoice

It is a total sum of a Charges for an Usage of a Service for an Account holder produced by a Provider
in the scope a Billing Event.

It optionally reference another Invoice being its correction by an UUID.

It has a rough structure like:

- Provider info
- Invoice number
- Customer info
  - Account
    - Charges
  - Total charges amount per account
- Invoice amount
- Tax amount
- Total invoice amount
- Invoice open amount (invoice amount may be paid with cash on account)

The structure will be stored as a JSON relational object.

### Ledger

An Charge or a Payment references an Ledger by UUID.

### Tax

It defines final tax to be calculated. Multiple taxes may be applied to an Invoice.

### Payment

A Payment may produce a Cash on account referencing an Invoice by UUID.

### Cash

A Cash references bot an Invoice and a Payment by their UUID.

## System components

This is just a rough conceptual model. The main requirement is that runs in K8S, it uses
a relational database to materialize entities and an off the shelf data bus
to source the events.

### Database

Any relational DB, Postgress shall be fine but considering Oracle as well. Less Golang
friendly as there is no ARM Oracle client yet.

The data model is created on the fly using the model entieties JSON structures.

### Data bus

The Data Bus is where an Event lives. It is created by an Actor and it is handled by a Hook.
Let's use Kafka.

### API

AN api provides CRUD actions to be performed on model entities by a controller. It shall have
both REST and/or gRPC services.

### Actor

An Actor is a process which can create an Event. It can produce an artefact as well.
It uses temporals to perform its actions. It polls on external input interfaces.
It may be triggered by intrnal input to produce an artefact in the output interface.

### Temporal

We can use temporal to implement an Actor. It offers an interface to process
handling.

Vide https://temporal.io/

### Hook

A Hook is an Event handler reacting to a particular even like a Billing request or Service
creation or activation (a type of Update).

It can invoke an API CRUD action but a Temporal Job as well.

### K8S

Let's do it in a K8S cluster. Docker, docker compose, controller to manage the system, etc.

### Deployment

In the cloud but very generic so that it can run in AWS, Azure, Google or any other
future possible cloud.