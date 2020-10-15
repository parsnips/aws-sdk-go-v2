// Code generated by smithy-go-codegen DO NOT EDIT.

// Package codestarnotifications provides the API client, operations, and parameter
// types for AWS CodeStar Notifications.
//
// This AWS CodeStar Notifications API Reference provides descriptions and usage
// examples of the operations and data types for the AWS CodeStar Notifications
// API. You can use the AWS CodeStar Notifications API to work with the following
// objects: Notification rules, by calling the following:
//
//     *
// CreateNotificationRule, which creates a notification rule for a resource in your
// account.
//
//     * DeleteNotificationRule, which deletes a notification rule.
//
//
// * DescribeNotificationRule, which provides information about a notification
// rule.
//
//     * ListNotificationRules, which lists the notification rules
// associated with your account.
//
//     * UpdateNotificationRule, which changes the
// name, events, or targets associated with a notification rule.
//
//     * Subscribe,
// which subscribes a target to a notification rule.
//
//     * Unsubscribe, which
// removes a target from a notification rule.
//
// Targets, by calling the following:
//
//
// * DeleteTarget, which removes a notification rule target (SNS topic) from a
// notification rule.
//
//     * ListTargets, which lists the targets associated with a
// notification rule.
//
// Events, by calling the following:
//
//     * ListEventTypes,
// which lists the event types you can include in a notification rule.
//
// Tags, by
// calling the following:
//
//     * ListTagsForResource, which lists the tags already
// associated with a notification rule in your account.
//
//     * TagResource, which
// associates a tag you provide with a notification rule in your account.
//
//     *
// UntagResource, which removes a tag from a notification rule in your
// account.
//
// For information about how to use AWS CodeStar Notifications, see link
// in the CodeStarNotifications User Guide.
package codestarnotifications
