# Yet another billing system

## Purpose

Starting with a list of main entieties to be included in a API for generic
billing system.

## The ontology

* Event

It is a potential change in the system state. It can be CRUD on any of the entieties
listed below.

* Service

It allows an Account holder to generate Usage to be chaerged by a Provider.

* Usage

It is an event of a USage of a Service by an Account.

* Fee

Is is a template for a Charge to be included in an Invoice,

* Subscription

It grants rights to an Account holder for usage of particular Service.

* Provider

It charges the holder of an Account for Usage of a Service

* Account

It is a representation of a physical holder of right to use a Service.

* Billing

It is a event of production of a Charge for a Service for an Account holder.

* Charge

It is an event of production of a Fee for a Usage of Service on an Account holder.

* Invoice (of a Charge for an Account)

It is a sum of charges for an Usage of a Service.
