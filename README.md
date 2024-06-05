# DEBILL

Also known as Durable Event Billing Interface Lazy Layer.

## Purpose

As usual, a PoC to note main entieties and infra components.

## The ontology

Starting with a list of main entieties to be included in a API for a very generic
billing system.

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

## System components

This is just a rough conceptual model. The main requirement is that runs in K8S, it uses
a relational database to materialize entities and an off the shelf data bus
to source the events.

### Database

Any relational DB. Postgress shall be fine but considering Oracle as well. Less Golang
friendly as there is no ARM Oracle client yet.

### Data bus

The Data Bus is where an Event lives. It is created by an Actor and it is handled by a Hook.
Let's use Kafka.

### API

AN api provides CRUD actions to be performed on model entities by a controller. It shall have
both REST and gRPC services.

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