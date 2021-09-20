package usecases

import (
	_activities "weatherit/usecases/activities"
	_alterplans "weatherit/usecases/alterplan"
	_events "weatherit/usecases/events"
	_interests "weatherit/usecases/interests"
	_userInterests "weatherit/usecases/user_interests"
	_users "weatherit/usecases/users"
)

type UsecaseList struct {
	Event        _events.UseCase
	Activity     _activities.UseCase
	AlterPlan    _alterplans.UseCase
	User         _users.UseCase
	Interest     _interests.UseCase
	UserInterest _userInterests.UseCase
}
