// Code generated by entc, DO NOT EDIT.

package model

import (
	"lin-cms-go/internal/data/ent/schema"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linlog"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/linuseridentiy"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	lingroupMixin := schema.LinGroup{}.Mixin()
	lingroupMixinFields0 := lingroupMixin[0].Fields()
	_ = lingroupMixinFields0
	lingroupFields := schema.LinGroup{}.Fields()
	_ = lingroupFields
	// lingroupDescCreateTime is the schema descriptor for create_time field.
	lingroupDescCreateTime := lingroupMixinFields0[0].Descriptor()
	// lingroup.DefaultCreateTime holds the default value on creation for the create_time field.
	lingroup.DefaultCreateTime = lingroupDescCreateTime.Default.(func() time.Time)
	// lingroupDescUpdateTime is the schema descriptor for update_time field.
	lingroupDescUpdateTime := lingroupMixinFields0[1].Descriptor()
	// lingroup.DefaultUpdateTime holds the default value on creation for the update_time field.
	lingroup.DefaultUpdateTime = lingroupDescUpdateTime.Default.(func() time.Time)
	// lingroup.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	lingroup.UpdateDefaultUpdateTime = lingroupDescUpdateTime.UpdateDefault.(func() time.Time)
	// lingroupDescDeleteTime is the schema descriptor for delete_time field.
	lingroupDescDeleteTime := lingroupMixinFields0[2].Descriptor()
	// lingroup.DefaultDeleteTime holds the default value on creation for the delete_time field.
	lingroup.DefaultDeleteTime = lingroupDescDeleteTime.Default.(func() time.Time)
	linlogMixin := schema.LinLog{}.Mixin()
	linlogMixinFields0 := linlogMixin[0].Fields()
	_ = linlogMixinFields0
	linlogFields := schema.LinLog{}.Fields()
	_ = linlogFields
	// linlogDescCreateTime is the schema descriptor for create_time field.
	linlogDescCreateTime := linlogMixinFields0[0].Descriptor()
	// linlog.DefaultCreateTime holds the default value on creation for the create_time field.
	linlog.DefaultCreateTime = linlogDescCreateTime.Default.(func() time.Time)
	// linlogDescUpdateTime is the schema descriptor for update_time field.
	linlogDescUpdateTime := linlogMixinFields0[1].Descriptor()
	// linlog.DefaultUpdateTime holds the default value on creation for the update_time field.
	linlog.DefaultUpdateTime = linlogDescUpdateTime.Default.(func() time.Time)
	// linlog.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	linlog.UpdateDefaultUpdateTime = linlogDescUpdateTime.UpdateDefault.(func() time.Time)
	// linlogDescDeleteTime is the schema descriptor for delete_time field.
	linlogDescDeleteTime := linlogMixinFields0[2].Descriptor()
	// linlog.DefaultDeleteTime holds the default value on creation for the delete_time field.
	linlog.DefaultDeleteTime = linlogDescDeleteTime.Default.(func() time.Time)
	linuserMixin := schema.LinUser{}.Mixin()
	linuserMixinFields0 := linuserMixin[0].Fields()
	_ = linuserMixinFields0
	linuserFields := schema.LinUser{}.Fields()
	_ = linuserFields
	// linuserDescCreateTime is the schema descriptor for create_time field.
	linuserDescCreateTime := linuserMixinFields0[0].Descriptor()
	// linuser.DefaultCreateTime holds the default value on creation for the create_time field.
	linuser.DefaultCreateTime = linuserDescCreateTime.Default.(func() time.Time)
	// linuserDescUpdateTime is the schema descriptor for update_time field.
	linuserDescUpdateTime := linuserMixinFields0[1].Descriptor()
	// linuser.DefaultUpdateTime holds the default value on creation for the update_time field.
	linuser.DefaultUpdateTime = linuserDescUpdateTime.Default.(func() time.Time)
	// linuser.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	linuser.UpdateDefaultUpdateTime = linuserDescUpdateTime.UpdateDefault.(func() time.Time)
	// linuserDescDeleteTime is the schema descriptor for delete_time field.
	linuserDescDeleteTime := linuserMixinFields0[2].Descriptor()
	// linuser.DefaultDeleteTime holds the default value on creation for the delete_time field.
	linuser.DefaultDeleteTime = linuserDescDeleteTime.Default.(func() time.Time)
	// linuserDescAvatar is the schema descriptor for avatar field.
	linuserDescAvatar := linuserFields[2].Descriptor()
	// linuser.DefaultAvatar holds the default value on creation for the avatar field.
	linuser.DefaultAvatar = linuserDescAvatar.Default.(string)
	linuseridentiyMixin := schema.LinUserIdentiy{}.Mixin()
	linuseridentiyMixinFields0 := linuseridentiyMixin[0].Fields()
	_ = linuseridentiyMixinFields0
	linuseridentiyFields := schema.LinUserIdentiy{}.Fields()
	_ = linuseridentiyFields
	// linuseridentiyDescCreateTime is the schema descriptor for create_time field.
	linuseridentiyDescCreateTime := linuseridentiyMixinFields0[0].Descriptor()
	// linuseridentiy.DefaultCreateTime holds the default value on creation for the create_time field.
	linuseridentiy.DefaultCreateTime = linuseridentiyDescCreateTime.Default.(func() time.Time)
	// linuseridentiyDescUpdateTime is the schema descriptor for update_time field.
	linuseridentiyDescUpdateTime := linuseridentiyMixinFields0[1].Descriptor()
	// linuseridentiy.DefaultUpdateTime holds the default value on creation for the update_time field.
	linuseridentiy.DefaultUpdateTime = linuseridentiyDescUpdateTime.Default.(func() time.Time)
	// linuseridentiy.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	linuseridentiy.UpdateDefaultUpdateTime = linuseridentiyDescUpdateTime.UpdateDefault.(func() time.Time)
	// linuseridentiyDescDeleteTime is the schema descriptor for delete_time field.
	linuseridentiyDescDeleteTime := linuseridentiyMixinFields0[2].Descriptor()
	// linuseridentiy.DefaultDeleteTime holds the default value on creation for the delete_time field.
	linuseridentiy.DefaultDeleteTime = linuseridentiyDescDeleteTime.Default.(func() time.Time)
}
