package goBungieNet

type errorCode int

const (
  ErrNone errorCode = 0
  ErrSuccess = 1
  ErrTransportException = 2
  ErrUnhandledException = 3
  ErrNotImplemented = 4
  ErrSystemDisabled = 5
  ErrFailedToLoadAvailableLocalesConfiguration = 6
  ErrParameterParseFailure = 7
  ErrParameterInvalidRange = 8
  ErrBadRequest = 9
  ErrAuthenticationInvalid = 10
  ErrDataNotFound = 11
  ErrInsufficientPrivileges = 12
  ErrDuplicate = 13
  ErrErrValidationError = 15
  ErrValidationMissingFieldError = 16
  ErrValidationInvalidInputError = 17
  ErrInvalidParameters = 18
  ErrParameterNotFound = 19
  ErrUnhandledHttpException = 20
  ErrNotFound = 21
  ErrWebAuthModuleAsyncFailed = 22
  ErrInvalidReturnValue = 23
  ErrUserBanned = 24
  ErrInvalidPostBody = 25
  ErrMissingPostBody = 26
  ErrExternalServiceTimeout = 27
  ErrValidationLengthError = 28
  ErrValidationRangeError = 29
  ErrJsonDeserializationError = 30
  ErrThrottleLimitExceeded = 31
  ErrValidationTagError = 32
  ErrValidationProfanityError = 33
  ErrValidationUrlFormatError = 34
  ErrThrottleLimitExceededMinutes = 35
  ErrThrottleLimitExceededMomentarily = 36
  ErrThrottleLimitExceededSeconds = 37
  ErrExternalServiceUnknown = 38
  ErrValidationWordLengthError = 39
  ErrValidationInvisibleUnicode = 40
  ErrValidationBadNames = 41
  ErrExternalServiceFailed = 42
  ErrServiceRetired = 43
  ErrUnknownSqlException = 44
  ErrUnsupportedLocale = 45
  ErrInvalidPageNumber = 46
  ErrMaximumPageSizeExceeded = 47
  ErrServiceUnsupported = 48
  ErrValidationMaximumUnicodeCombiningCharacters = 49
  ErrValidationMaximumSequentialCarriageReturns = 50
  ErrPerEndpointRequestThrottleExceeded = 51
  ErrAuthContextCacheAssertion = 52
  ErrObsoleteCredentialType = 89
  ErrUnableToUnPairMobileApp = 90
  ErrUnableToPairMobileApp = 91
  ErrCannotUseMobileAuthWithNonMobileProvider = 92
  ErrMissingDeviceCookie = 93
  ErrFacebookTokenExpired = 94
  ErrAuthTicketRequired = 95
  ErrCookieContextRequired = 96
  ErrUnknownAuthenticationError = 97
  ErrBungieNetAccountCreationRequired = 98
  ErrWebAuthRequired = 99
  ErrContentUnknownSqlResult = 100
  ErrContentNeedUniquePath = 101
  ErrContentSqlException = 102
  ErrContentNotFound = 103
  ErrContentSuccessWithTagAddFail = 104
  ErrContentSearchMissingParameters = 105
  ErrContentInvalidId = 106
  ErrContentPhysicalFileDeletionError = 107
  ErrContentPhysicalFileCreationError = 108
  ErrContentPerforceSubmissionError = 109
  ErrContentPerforceInitializationError = 110
  ErrContentDeploymentPackageNotReadyError = 111
  ErrContentUploadFailed = 112
  ErrContentTooManyResults = 113
  ErrContentInvalidState = 115
  ErrContentNavigationParentNotFound = 116
  ErrContentNavigationParentUpdateError = 117
  ErrDeploymentPackageNotEditable = 118
  ErrContentValidationError = 119
  ErrContentPropertiesValidationError = 120
  ErrContentTypeNotFound = 121
  ErrDeploymentPackageNotFound = 122
  ErrContentSearchInvalidParameters = 123
  ErrContentItemPropertyAggregationError = 124
  ErrDeploymentPackageFileNotFound = 125
  ErrContentPerforceFileHistoryNotFound = 126
  ErrContentAssetZipCreationFailure = 127
  ErrContentAssetZipCreationBusy = 128
  ErrContentProjectNotFound = 129
  ErrContentFolderNotFound = 130
  ErrContentPackagesInconsistent = 131
  ErrContentPackagesInvalidState = 132
  ErrContentPackagesInconsistentType = 133
  ErrContentCannotDeletePackage = 134
  ErrContentLockedForChanges = 135
  ErrContentFileUploadFailed = 136
  ErrContentNotReviewed = 137
  ErrContentPermissionDenied = 138
  ErrContentInvalidExternalUrl = 139
  ErrContentExternalFileCannotBeImportedLocally = 140
  ErrContentTagSaveFailure = 141
  ErrContentPerforceUnmatchedFileError = 142
  ErrContentPerforceChangelistResultNotFound = 143
  ErrContentPerforceChangelistFileItemsNotFound = 144
  ErrContentPerforceInvalidRevisionError = 145
  ErrContentUnloadedSaveResult = 146
  ErrContentPropertyInvalidNumber = 147
  ErrContentPropertyInvalidUrl = 148
  ErrContentPropertyInvalidDate = 149
  ErrContentPropertyInvalidSet = 150
  ErrContentPropertyCannotDeserialize = 151
  ErrContentRegexValidationFailOnProperty = 152
  ErrContentMaxLengthFailOnProperty = 153
  ErrContentPropertyUnexpectedDeserializationError = 154
  ErrContentPropertyRequired = 155
  ErrContentCannotCreateFile = 156
  ErrContentInvalidMigrationFile = 157
  ErrContentMigrationAlteringProcessedItem = 158
  ErrContentPropertyDefinitionNotFound = 159
  ErrContentReviewDataChanged = 160
  ErrContentRollbackRevisionNotInPackage = 161
  ErrContentItemNotBasedOnLatestRevision = 162
  ErrContentUnauthorized = 163
  ErrContentCannotCreateDeploymentPackage = 164
  ErrContentUserNotFound = 165
  ErrContentLocalePermissionDenied = 166
  ErrContentInvalidLinkToInternalEnvironment = 167
  ErrContentInvalidBlacklistedContent = 168
  ErrContentMacroMalformedNoContentId = 169
  ErrContentMacroMalformedNoTemplateType = 170
  ErrContentIllegalBNetMembershipId = 171
  ErrContentLocaleDidNotMatchExpected = 172
  ErrContentBabelCallFailed = 173
  ErrContentEnglishPostLiveForbidden = 174
  ErrContentLocaleEditPermissionDenied = 175
  ErrUserNonUniqueName = 200
  ErrUserManualLinkingStepRequired = 201
  ErrUserCreateUnknownSqlResult = 202
  ErrUserCreateUnknownSqlException = 203
  ErrUserMalformedMembershipId = 204
  ErrUserCannotFindRequestedUser = 205
  ErrUserCannotLoadAccountCredentialLinkInfo = 206
  ErrUserInvalidMobileAppType = 207
  ErrUserMissingMobilePairingInfo = 208
  ErrUserCannotGenerateMobileKeyWhileUsingMobileCredential = 209
  ErrUserGenerateMobileKeyExistingSlotCollision = 210
  ErrUserDisplayNameMissingOrInvalid = 211
  ErrUserCannotLoadAccountProfileData = 212
  ErrUserCannotSaveUserProfileData = 213
  ErrUserEmailMissingOrInvalid = 214
  ErrUserTermsOfUseRequired = 215
  ErrUserCannotCreateNewAccountWhileLoggedIn = 216
  ErrUserCannotResolveCentralAccount = 217
  ErrUserInvalidAvatar = 218
  ErrUserMissingCreatedUserResult = 219
  ErrUserCannotChangeUniqueNameYet = 220
  ErrUserCannotChangeDisplayNameYet = 221
  ErrUserCannotChangeEmail = 222
  ErrUserUniqueNameMustStartWithLetter = 223
  ErrUserNoLinkedAccountsSupportFriendListings = 224
  ErrUserAcknowledgmentTableFull = 225
  ErrUserCreationDestinyMembershipRequired = 226
  ErrUserFriendsTokenNeedsRefresh = 227
  ErrMessagingUnknownError = 300
  ErrMessagingSelfError = 301
  ErrMessagingSendThrottle = 302
  ErrMessagingNoBody = 303
  ErrMessagingTooManyUsers = 304
  ErrMessagingCanNotLeaveConversation = 305
  ErrMessagingUnableToSend = 306
  ErrMessagingDeletedUserForbidden = 307
  ErrMessagingCannotDeleteExternalConversation = 308
  ErrMessagingGroupChatDisabled = 309
  ErrMessagingMustIncludeSelfInPrivateMessage = 310
  ErrMessagingSenderIsBanned = 311
  ErrMessagingGroupOptionalChatExceededMaximum = 312
  ErrAddSurveyAnswersUnknownSqlException = 400
  ErrForumBodyCannotBeEmpty = 500
  ErrForumSubjectCannotBeEmptyOnTopicPost = 501
  ErrForumCannotLocateParentPost = 502
  ErrForumThreadLockedForReplies = 503
  ErrForumUnknownSqlResultDuringCreatePost = 504
  ErrForumUnknownTagCreationError = 505
  ErrForumUnknownSqlResultDuringTagItem = 506
  ErrForumUnknownExceptionCreatePost = 507
  ErrForumQuestionMustBeTopicPost = 508
  ErrForumExceptionDuringTagSearch = 509
  ErrForumExceptionDuringTopicRetrieval = 510
  ErrForumAliasedTagError = 511
  ErrForumCannotLocateThread = 512
  ErrForumUnknownExceptionEditPost = 513
  ErrForumCannotLocatePost = 514
  ErrForumUnknownExceptionGetOrCreateTags = 515
  ErrForumEditPermissionDenied = 516
  ErrForumUnknownSqlResultDuringTagIdRetrieval = 517
  ErrForumCannotGetRating = 518
  ErrForumUnknownExceptionGetRating = 519
  ErrForumRatingsAccessError = 520
  ErrForumRelatedPostAccessError = 521
  ErrForumLatestReplyAccessError = 522
  ErrForumUserStatusAccessError = 523
  ErrForumAuthorAccessError = 524
  ErrForumGroupAccessError = 525
  ErrForumUrlExpectedButMissing = 526
  ErrForumRepliesCannotBeEmpty = 527
  ErrForumRepliesCannotBeInDifferentGroups = 528
  ErrForumSubTopicCannotBeCreatedAtThisThreadLevel = 529
  ErrForumCannotCreateContentTopic = 530
  ErrForumTopicDoesNotExist = 531
  ErrForumContentCommentsNotAllowed = 532
  ErrForumUnknownSqlResultDuringEditPost = 533
  ErrForumUnknownSqlResultDuringGetPost = 534
  ErrForumPostValidationBadUrl = 535
  ErrForumBodyTooLong = 536
  ErrForumSubjectTooLong = 537
  ErrForumAnnouncementNotAllowed = 538
  ErrForumCannotShareOwnPost = 539
  ErrForumEditNoOp = 540
  ErrForumUnknownDatabaseErrorDuringGetPost = 541
  ErrForumExceeedMaximumRowLimit = 542
  ErrForumCannotSharePrivatePost = 543
  ErrForumCannotCrossPostBetweenGroups = 544
  ErrForumIncompatibleCategories = 555
  ErrForumCannotUseTheseCategoriesOnNonTopicPost = 556
  ErrForumCanOnlyDeleteTopics = 557
  ErrForumDeleteSqlException = 558
  ErrForumDeleteSqlUnknownResult = 559
  ErrForumTooManyTags = 560
  ErrForumCanOnlyRateTopics = 561
  ErrForumBannedPostsCannotBeEdited = 562
  ErrForumThreadRootIsBanned = 563
  ErrForumCannotUseOfficialTagCategoryAsTag = 564
  ErrForumAnswerCannotBeMadeOnCreatePost = 565
  ErrForumAnswerCannotBeMadeOnEditPost = 566
  ErrForumAnswerPostIdIsNotADirectReplyOfQuestion = 567
  ErrForumAnswerTopicIdIsNotAQuestion = 568
  ErrForumUnknownExceptionDuringMarkAnswer = 569
  ErrForumUnknownSqlResultDuringMarkAnswer = 570
  ErrForumCannotRateYourOwnPosts = 571
  ErrForumPollsMustBeTheFirstPostInTopic = 572
  ErrForumInvalidPollInput = 573
  ErrForumGroupAdminEditNonMember = 574
  ErrForumCannotEditModeratorEditedPost = 575
  ErrForumRequiresDestinyMembership = 576
  ErrForumUnexpectedError = 577
  ErrForumAgeLock = 578
  ErrForumMaxPages = 579
  ErrForumMaxPagesOldestFirst = 580
  ErrForumCannotApplyForumIdWithoutTags = 581
  ErrForumCannotApplyForumIdToNonTopics = 582
  ErrForumCannotDownvoteCommunityCreations = 583
  ErrForumTopicsMustHaveOfficialCategory = 584
  ErrForumRecruitmentTopicMalformed = 585
  ErrForumRecruitmentTopicNotFound = 586
  ErrForumRecruitmentTopicNoSlotsRemaining = 587
  ErrForumRecruitmentTopicKickBan = 588
  ErrForumRecruitmentTopicRequirementsNotMet = 589
  ErrForumRecruitmentTopicNoPlayers = 590
  ErrForumRecruitmentApproveFailMessageBan = 591
  ErrForumRecruitmentGlobalBan = 592
  ErrForumUserBannedFromThisTopic = 593
  ErrForumRecruitmentFireteamMembersOnly = 594
  ErrGroupMembershipApplicationAlreadyResolved = 601
  ErrGroupMembershipAlreadyApplied = 602
  ErrGroupMembershipInsufficientPrivileges = 603
  ErrGroupIdNotReturnedFromCreation = 604
  ErrGroupSearchInvalidParameters = 605
  ErrGroupMembershipPendingApplicationNotFound = 606
  ErrGroupInvalidId = 607
  ErrGroupInvalidMembershipId = 608
  ErrGroupInvalidMembershipType = 609
  ErrGroupMissingTags = 610
  ErrGroupMembershipNotFound = 611
  ErrGroupInvalidRating = 612
  ErrGroupUserFollowingAccessError = 613
  ErrGroupUserMembershipAccessError = 614
  ErrGroupCreatorAccessError = 615
  ErrGroupAdminAccessError = 616
  ErrGroupPrivatePostNotViewable = 617
  ErrGroupMembershipNotLoggedIn = 618
  ErrGroupNotDeleted = 619
  ErrGroupUnknownErrorUndeletingGroup = 620
  ErrGroupDeleted = 621
  ErrGroupNotFound = 622
  ErrGroupMemberBanned = 623
  ErrGroupMembershipClosed = 624
  ErrGroupPrivatePostOverrideError = 625
  ErrGroupNameTaken = 626
  ErrGroupDeletionGracePeriodExpired = 627
  ErrGroupCannotCheckBanStatus = 628
  ErrGroupMaximumMembershipCountReached = 629
  ErrNoDestinyAccountForClanPlatform = 630
  ErrAlreadyRequestingMembershipForClanPlatform = 631
  ErrAlreadyClanMemberOnPlatform = 632
  ErrGroupJoinedCannotSetClanName = 633
  ErrGroupLeftCannotClearClanName = 634
  ErrGroupRelationshipRequestPending = 635
  ErrGroupRelationshipRequestBlocked = 636
  ErrGroupRelationshipRequestNotFound = 637
  ErrGroupRelationshipBlockNotFound = 638
  ErrGroupRelationshipNotFound = 639
  ErrGroupAlreadyAllied = 641
  ErrGroupAlreadyMember = 642
  ErrGroupRelationshipAlreadyExists = 643
  ErrInvalidGroupTypesForRelationshipRequest = 644
  ErrGroupAtMaximumAlliances = 646
  ErrGroupCannotSetClanOnlySettings = 647
  ErrClanCannotSetTwoDefaultPostTypes = 648
  ErrGroupMemberInvalidMemberType = 649
  ErrGroupInvalidPlatformType = 650
  ErrGroupMemberInvalidSort = 651
  ErrGroupInvalidResolveState = 652
  ErrClanAlreadyEnabledForPlatform = 653
  ErrClanNotEnabledForPlatform = 654
  ErrClanEnabledButCouldNotJoinNoAccount = 655
  ErrClanEnabledButCouldNotJoinAlreadyMember = 656
  ErrClanCannotJoinNoCredential = 657
  ErrNoClanMembershipForPlatform = 658
  ErrGroupToGroupFollowLimitReached = 659
  ErrChildGroupAlreadyInAlliance = 660
  ErrOwnerGroupAlreadyInAlliance = 661
  ErrAllianceOwnerCannotJoinAlliance = 662
  ErrGroupNotInAlliance = 663
  ErrChildGroupCannotInviteToAlliance = 664
  ErrGroupToGroupAlreadyFollowed = 665
  ErrGroupToGroupNotFollowing = 666
  ErrClanMaximumMembershipReached = 667
  ErrClanNameNotValid = 668
  ErrClanNameNotValidError = 669
  ErrAllianceOwnerNotDefined = 670
  ErrAllianceChildNotDefined = 671
  ErrClanCultureIllegalCharacters = 672
  ErrClanTagIllegalCharacters = 673
  ErrClanRequiresInvitation = 674
  ErrClanMembershipClosed = 675
  ErrClanInviteAlreadyMember = 676
  ErrGroupInviteAlreadyMember = 677
  ErrGroupJoinApprovalRequired = 678
  ErrClanTagRequired = 679
  ErrGroupNameCannotStartOrEndWithWhiteSpace = 680
  ErrClanCallsignCannotStartOrEndWithWhiteSpace = 681
  ErrClanMigrationFailed = 682
  ErrClanNotEnabledAlreadyMemberOfAnotherClan = 683
  ErrGroupModerationNotPermittedOnNonMembers = 684
  ErrClanCreationInWorldServerFailed = 685
  ErrClanNotFound = 686
  ErrClanMembershipLevelDoesNotPermitThatAction = 687
  ErrClanMemberNotFound = 688
  ErrClanMissingMembershipApprovers = 689
  ErrClanInWrongStateForRequestedAction = 690
  ErrClanNameAlreadyUsed = 691
  ErrClanTooFewMembers = 692
  ErrClanInfoCannotBeWhitespace = 693
  ErrGroupCultureThrottle = 694
  ErrClanTargetDisallowsInvites = 695
  ErrClanInvalidOperation = 696
  ErrClanFounderCannotLeaveWithoutAbdication = 697
  ErrClanNameReserved = 698
  ErrClanApplicantInClanSoNowInvited = 699
  ErrActivitiesUnknownException = 701
  ErrActivitiesParameterNull = 702
  ErrActivityCountsDiabled = 703
  ErrActivitySearchInvalidParameters = 704
  ErrActivityPermissionDenied = 705
  ErrShareAlreadyShared = 706
  ErrActivityLoggingDisabled = 707
  ErrItemAlreadyFollowed = 801
  ErrItemNotFollowed = 802
  ErrCannotFollowSelf = 803
  ErrGroupFollowLimitExceeded = 804
  ErrTagFollowLimitExceeded = 805
  ErrUserFollowLimitExceeded = 806
  ErrFollowUnsupportedEntityType = 807
  ErrNoValidTagsInList = 900
  ErrBelowMinimumSuggestionLength = 901
  ErrCannotGetSuggestionsOnMultipleTagsSimultaneously = 902
  ErrNotAValidPartialTag = 903
  ErrTagSuggestionsUnknownSqlResult = 904
  ErrTagsUnableToLoadPopularTagsFromDatabase = 905
  ErrTagInvalid = 906
  ErrTagNotFound = 907
  ErrSingleTagExpected = 908
  ErrTagsExceededMaximumPerItem = 909
  ErrIgnoreInvalidParameters = 1000
  ErrIgnoreSqlException = 1001
  ErrIgnoreErrorRetrievingGroupPermissions = 1002
  ErrIgnoreErrorInsufficientPermission = 1003
  ErrIgnoreErrorRetrievingItem = 1004
  ErrIgnoreCannotIgnoreSelf = 1005
  ErrIgnoreIllegalType = 1006
  ErrIgnoreNotFound = 1007
  ErrIgnoreUserGloballyIgnored = 1008
  ErrIgnoreUserIgnored = 1009
  ErrNotificationSettingInvalid = 1100
  ErrPsnApiExpiredAccessToken = 1204
  ErrPSNExForbidden = 1205
  ErrPSNExSystemDisabled = 1218
  ErrPsnApiErrorCodeUnknown = 1223
  ErrPsnApiErrorWebException = 1224
  ErrPsnApiBadRequest = 1225
  ErrPsnApiAccessTokenRequired = 1226
  ErrPsnApiInvalidAccessToken = 1227
  ErrPsnApiBannedUser = 1229
  ErrPsnApiAccountUpgradeRequired = 1230
  ErrPsnApiServiceTemporarilyUnavailable = 1231
  ErrPsnApiServerBusy = 1232
  ErrPsnApiUnderMaintenance = 1233
  ErrPsnApiProfileUserNotFound = 1234
  ErrPsnApiProfilePrivacyRestriction = 1235
  ErrPsnApiProfileUnderMaintenance = 1236
  ErrPsnApiAccountAttributeMissing = 1237
  ErrXblExSystemDisabled = 1300
  ErrXblExUnknownError = 1301
  ErrXblApiErrorWebException = 1302
  ErrXblStsTokenInvalid = 1303
  ErrXblStsMissingToken = 1304
  ErrXblStsExpiredToken = 1305
  ErrXblAccessToTheSandboxDenied = 1306
  ErrXblMsaResponseMissing = 1307
  ErrXblMsaAccessTokenExpired = 1308
  ErrXblMsaInvalidRequest = 1309
  ErrXblMsaFriendsRequireSignIn = 1310
  ErrXblUserActionRequired = 1311
  ErrXblParentalControls = 1312
  ErrXblDeveloperAccount = 1313
  ErrXblUserTokenExpired = 1314
  ErrXblUserTokenInvalid = 1315
  ErrXblOffline = 1316
  ErrXblUnknownErrorCode = 1317
  ErrXblMsaInvalidGrant = 1318
  ErrReportNotYetResolved = 1400
  ErrReportOverturnDoesNotChangeDecision = 1401
  ErrReportNotFound = 1402
  ErrReportAlreadyReported = 1403
  ErrReportInvalidResolution = 1404
  ErrReportNotAssignedToYou = 1405
  ErrLegacyGameStatsSystemDisabled = 1500
  ErrLegacyGameStatsUnknownError = 1501
  ErrLegacyGameStatsMalformedSneakerNetCode = 1502
  ErrDestinyAccountAcquisitionFailure = 1600
  ErrDestinyAccountNotFound = 1601
  ErrDestinyBuildStatsDatabaseError = 1602
  ErrDestinyCharacterStatsDatabaseError = 1603
  ErrDestinyPvPStatsDatabaseError = 1604
  ErrDestinyPvEStatsDatabaseError = 1605
  ErrDestinyGrimoireStatsDatabaseError = 1606
  ErrDestinyStatsParameterMembershipTypeParseError = 1607
  ErrDestinyStatsParameterMembershipIdParseError = 1608
  ErrDestinyStatsParameterRangeParseError = 1609
  ErrDestinyStringItemHashNotFound = 1610
  ErrDestinyStringSetNotFound = 1611
  ErrDestinyContentLookupNotFoundForKey = 1612
  ErrDestinyContentItemNotFound = 1613
  ErrDestinyContentSectionNotFound = 1614
  ErrDestinyContentPropertyNotFound = 1615
  ErrDestinyContentConfigNotFound = 1616
  ErrDestinyContentPropertyBucketValueNotFound = 1617
  ErrDestinyUnexpectedError = 1618
  ErrDestinyInvalidAction = 1619
  ErrDestinyCharacterNotFound = 1620
  ErrDestinyInvalidFlag = 1621
  ErrDestinyInvalidRequest = 1622
  ErrDestinyItemNotFound = 1623
  ErrDestinyInvalidCustomizationChoices = 1624
  ErrDestinyVendorItemNotFound = 1625
  ErrDestinyInternalError = 1626
  ErrDestinyVendorNotFound = 1627
  ErrDestinyRecentActivitiesDatabaseError = 1628
  ErrDestinyItemBucketNotFound = 1629
  ErrDestinyInvalidMembershipType = 1630
  ErrDestinyVersionIncompatibility = 1631
  ErrDestinyItemAlreadyInInventory = 1632
  ErrDestinyBucketNotFound = 1633
  ErrDestinyCharacterNotInTower = 1634
  ErrDestinyCharacterNotLoggedIn = 1635
  ErrDestinyDefinitionsNotLoaded = 1636
  ErrDestinyInventoryFull = 1637
  ErrDestinyItemFailedLevelCheck = 1638
  ErrDestinyItemFailedUnlockCheck = 1639
  ErrDestinyItemUnequippable = 1640
  ErrDestinyItemUniqueEquipRestricted = 1641
  ErrDestinyNoRoomInDestination = 1642
  ErrDestinyServiceFailure = 1643
  ErrDestinyServiceRetired = 1644
  ErrDestinyTransferFailed = 1645
  ErrDestinyTransferNotFoundForSourceBucket = 1646
  ErrDestinyUnexpectedResultInVendorTransferCheck = 1647
  ErrDestinyUniquenessViolation = 1648
  ErrDestinyErrorDeserializationFailure = 1649
  ErrDestinyValidAccountTicketRequired = 1650
  ErrDestinyShardRelayClientTimeout = 1651
  ErrDestinyShardRelayProxyTimeout = 1652
  ErrDestinyPGCRNotFound = 1653
  ErrDestinyAccountMustBeOffline = 1654
  ErrDestinyCanOnlyEquipInGame = 1655
  ErrDestinyCannotPerformActionOnEquippedItem = 1656
  ErrDestinyQuestAlreadyCompleted = 1657
  ErrDestinyQuestAlreadyTracked = 1658
  ErrDestinyTrackableQuestsFull = 1659
  ErrDestinyItemNotTransferrable = 1660
  ErrDestinyVendorPurchaseNotAllowed = 1661
  ErrDestinyContentVersionMismatch = 1662
  ErrDestinyItemActionForbidden = 1663
  ErrDestinyRefundInvalid = 1664
  ErrDestinyPrivacyRestriction = 1665
  ErrDestinyActionInsufficientPrivileges = 1666
  ErrDestinyInvalidClaimException = 1667
  ErrDestinyLegacyPlatformRestricted = 1668
  ErrDestinyLegacyPlatformInUse = 1669
  ErrDestinyLegacyPlatformInaccessible = 1670
  ErrDestinyCannotPerformActionAtThisLocation = 1671
  ErrDestinyThrottledByGameServer = 1672
  ErrFbInvalidRequest = 1800
  ErrFbRedirectMismatch = 1801
  ErrFbAccessDenied = 1802
  ErrFbUnsupportedResponseType = 1803
  ErrFbInvalidScope = 1804
  ErrFbUnsupportedGrantType = 1805
  ErrFbInvalidGrant = 1806
  ErrInvitationExpired = 1900
  ErrInvitationUnknownType = 1901
  ErrInvitationInvalidResponseStatus = 1902
  ErrInvitationInvalidType = 1903
  ErrInvitationAlreadyPending = 1904
  ErrInvitationInsufficientPermission = 1905
  ErrInvitationInvalidCode = 1906
  ErrInvitationInvalidTargetState = 1907
  ErrInvitationCannotBeReactivated = 1908
  ErrInvitationNoRecipients = 1910
  ErrInvitationGroupCannotSendToSelf = 1911
  ErrInvitationTooManyRecipients = 1912
  ErrInvitationInvalid = 1913
  ErrInvitationNotFound = 1914
  ErrTokenInvalid = 2000
  ErrTokenBadFormat = 2001
  ErrTokenAlreadyClaimed = 2002
  ErrTokenAlreadyClaimedSelf = 2003
  ErrTokenThrottling = 2004
  ErrTokenUnknownRedemptionFailure = 2005
  ErrTokenPurchaseClaimFailedAfterTokenClaimed = 2006
  ErrTokenUserAlreadyOwnsOffer = 2007
  ErrTokenInvalidOfferKey = 2008
  ErrTokenEmailNotValidated = 2009
  ErrTokenProvisioningBadVendorOrOffer = 2010
  ErrTokenPurchaseHistoryUnknownError = 2011
  ErrTokenThrottleStateUnknownError = 2012
  ErrTokenUserAgeNotVerified = 2013
  ErrTokenExceededOfferMaximum = 2014
  ErrTokenNoAvailableUnlocks = 2015
  ErrTokenMarketplaceInvalidPlatform = 2016
  ErrTokenNoMarketplaceCodesFound = 2017
  ErrTokenOfferNotAvailableForRedemption = 2018
  ErrTokenUnlockPartialFailure = 2019
  ErrTokenMarketplaceInvalidRegion = 2020
  ErrTokenOfferExpired = 2021
  ErrRAFExceededMaximumReferrals = 2022
  ErrRAFDuplicateBond = 2023
  ErrRAFNoValidVeteranDestinyMembershipsFound = 2024
  ErrRAFNotAValidVeteranUser = 2025
  ErrRAFCodeAlreadyClaimedOrNotFound = 2026
  ErrRAFMismatchedDestinyMembershipType = 2027
  ErrRAFUnableToAccessPurchaseHistory = 2028
  ErrRAFUnableToCreateBond = 2029
  ErrRAFUnableToFindBond = 2030
  ErrRAFUnableToRemoveBond = 2031
  ErrRAFCannotBondToSelf = 2032
  ErrRAFInvalidPlatform = 2033
  ErrRAFGenerateThrottled = 2034
  ErrRAFUnableToCreateBondVersionMismatch = 2035
  ErrRAFUnableToRemoveBondVersionMismatch = 2036
  ErrRAFRedeemThrottled = 2037
  ErrNoAvailableDiscountCode = 2038
  ErrDiscountAlreadyClaimed = 2039
  ErrDiscountClaimFailure = 2040
  ErrDiscountConfigurationFailure = 2041
  ErrDiscountGenerationFailure = 2042
  ErrDiscountAlreadyExists = 2043
  ErrTokenRequiresCredentialXuid = 2044
  ErrTokenRequiresCredentialPsnid = 2045
  ErrOfferRequired = 2046
  ErrApiExceededMaxKeys = 2100
  ErrApiInvalidOrExpiredKey = 2101
  ErrApiKeyMissingFromRequest = 2102
  ErrApplicationDisabled = 2103
  ErrApplicationExceededMax = 2104
  ErrApplicationDisallowedByScope = 2105
  ErrAuthorizationCodeInvalid = 2106
  ErrOriginHeaderDoesNotMatchKey = 2107
  ErrAccessNotPermittedByApplicationScope = 2108
  ErrApplicationNameIsTaken = 2109
  ErrRefreshTokenNotYetValid = 2110
  ErrAccessTokenHasExpired = 2111
  ErrApplicationTokenFormatNotValid = 2112
  ErrApplicationNotConfiguredForBungieAuth = 2113
  ErrApplicationNotConfiguredForOAuth = 2114
  ErrOAuthAccessTokenExpired = 2115
  ErrPartnershipInvalidType = 2200
  ErrPartnershipValidationError = 2201
  ErrPartnershipValidationTimeout = 2202
  ErrPartnershipAccessFailure = 2203
  ErrPartnershipAccountInvalid = 2204
  ErrPartnershipGetAccountInfoFailure = 2205
  ErrPartnershipDisabled = 2206
  ErrPartnershipAlreadyExists = 2207
  ErrCommunityStreamingUnavailable = 2300
  ErrTwitchNotLinked = 2500
  ErrTwitchAccountNotFound = 2501
  ErrTwitchCouldNotLoadDestinyInfo = 2502
  ErrTrendingCategoryNotFound = 2600
  ErrTrendingEntryTypeNotSupported = 2601
)

func (code errorCode)isError() bool {
  return code != ErrNone && code != ErrSuccess
}
