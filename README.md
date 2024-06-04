# Yet another billing system

## Purpose

The ontology of a general purpose billing system.

### Eventual entities

#### Service

It allows an Account holder to generate Usage.

#### Usage

It is an record of a Service being used by an Account.

#### Subscription

It grants rights to an Account holder for usage of particular Service.

#### Provider

It charges the holder of an Account for Usage of a Service

#### Account

It is a representation of a physical holder of right to use a Service.

#### Billing

It is a event of production of a Charge for a Service for an Account holder.

#### Charge

It is an event of production of fee for a Usage of Service on an Account holder.

#### Invoice (of a Charge for an Account)

It is a sum of charges for an Usage of a Service.
