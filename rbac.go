package gorbac

import "fmt"

type Operator string

const (
	EQUALS              Operator = "eq"
	NOT_EQUALS          Operator = "ne"
	GREATER_THAN_EQUALS Operator = "gte"
	GREATER_THAN        Operator = "gt"
	LESS_THAN           Operator = "lt"
	LESS_THAN_EQUALS    Operator = "lte"
	CONTAINS            Operator = "like"
	ICONTAINS           Operator = "ilike"
	ANY                 Operator = "any"
	IN                  Operator = "in"
)

type Rule struct {
	Allowed  bool        `json:"allowed" yaml:"allowed"`
	Key      string      `json:"key" yaml:"key"`
	Operator Operator    `json:"operator" yaml:"operator"`
	Value    interface{} `json:"value" yaml:"value"`
}

type Permission struct {
	ID     string `json:"id" yaml:"id"`
	Name   string `json:"name" yaml:"name"`
	Action string `json:"action" yaml:"action"`
	Target string `json:"target" yaml:"target"`
	Rules  []Rule `json:"rules" yaml:"rules"`
}

type Role struct {
	ID          string       `json:"id" yaml:"id"`
	Name        string       `json:"name" yaml:"name"`
	Permissions []Permission `json:"permissions" yaml:"permissions"`
}

type Rbac struct {
	roles []Role `json:"roles" yaml:"roles"`
}

type Action string

const (
	CREATE Action = "create"
	READ   Action = "read"
	UPDATE Action = "update"
	DELETE Action = "delete"
	PLAY   Action = "play"
	ADMIN  Action = "admin"
)

func With(roles []Role) *Rbac {
	return &Rbac{roles}
}

func (r *Rbac) HasPermission(action Action, target string) bool {
	for _, role := range r.roles {
		for _, permission := range role.Permissions {
			hasAction := fmt.Sprintf("%s", action) == permission.Action || fmt.Sprintf("%s", ADMIN) == permission.Action
			hasTarget := fmt.Sprintf("%s", target) == permission.Target || "*" == permission.Target

			if hasAction && hasTarget {
				return true
			}
		}
	}
	return false
}

func (r *Rbac) IsAllowed(i interface{}) bool {
	hashmap := Struct(i).ToMap()
	for _, role := range r.roles {
		for _, permission := range role.Permissions {
			isAllowed := r.verifyIsAllowed(hashmap, permission)
			if isAllowed {
				return isAllowed
			}
		}
	}
	return false
}

func (r *Rbac) verifyIsAllowed(hashmap map[string]interface{}, permission Permission) bool {
	for _, rule := range permission.Rules {
		if value, ok := hashmap[rule.Key]; ok {
			if !rule.Allowed {
				return false
			}
			switch value.(type) {
			case string:
				return r.checkRuleValueForString(rule, value.(string))
			case int, int32, int64, float32, float64:
				return r.checkRuleValueForNumber(rule, value.(int))
			case bool:
				return r.checkRuleValueForBool(rule, value.(bool))
			default:
				return false
			}
		}
	}
	return false
}

func (r *Rbac) checkRuleValueForString(rule Rule, val string) bool {
	switch rule.Value.(type) {
	case string:
		value := rule.Value.(string)
		switch rule.Operator {
		case EQUALS:
			return value == val
		case NOT_EQUALS:
			return value != val
		default:
			return false
		}
	default:
		return false
	}
}

func (r *Rbac) checkRuleValueForNumber(rule Rule, val int) bool {
	switch rule.Value.(type) {
	case int:
		value := rule.Value.(int)
		switch rule.Operator {
		case EQUALS:
			return value == val
		case NOT_EQUALS:
			return value != val
		case GREATER_THAN_EQUALS:
			return value >= val
		case GREATER_THAN:
			return value > val
		case LESS_THAN:
			return value < val
		case LESS_THAN_EQUALS:
			return value <= val
		default:
			return false
		}
	default:
		return false
	}
}

func (r *Rbac) checkRuleValueForBool(rule Rule, val bool) bool {
	switch rule.Value.(type) {
	case bool:
		value := rule.Value.(bool)
		switch rule.Operator {
		case EQUALS:
			return value == val
		case NOT_EQUALS:
			return value != val
		default:
			return false
		}
	default:
		return false
	}
}
