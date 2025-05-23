package console

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocksconsole "github.com/goravel/framework/mocks/console"
	"github.com/goravel/framework/support/file"
)

func TestFactoryMakeCommand(t *testing.T) {
	factoryMakeCommand := &FactoryMakeCommand{}
	mockContext := mocksconsole.NewContext(t)
	mockContext.EXPECT().Argument(0).Return("").Once()
	mockContext.EXPECT().Ask("Enter the factory name", mock.Anything).Return("", errors.New("the factory name cannot be empty")).Once()
	mockContext.EXPECT().Error("the factory name cannot be empty").Once()
	assert.NoError(t, factoryMakeCommand.Handle(mockContext))

	mockContext.EXPECT().Argument(0).Return("UserFactory").Once()
	mockContext.EXPECT().OptionBool("force").Return(false).Once()
	mockContext.EXPECT().Success("Factory created successfully").Once()
	assert.NoError(t, factoryMakeCommand.Handle(mockContext))
	assert.True(t, file.Exists("database/factories/user_factory.go"))
	assert.True(t, file.Contain("database/factories/user_factory.go", "package factories"))
	assert.True(t, file.Contain("database/factories/user_factory.go", "type UserFactory struct"))

	mockContext.EXPECT().Argument(0).Return("UserFactory").Once()
	mockContext.EXPECT().OptionBool("force").Return(false).Once()
	mockContext.EXPECT().Error("the factory already exists. Use the --force or -f flag to overwrite").Once()
	assert.NoError(t, factoryMakeCommand.Handle(mockContext))
	assert.NoError(t, file.Remove("database"))

	mockContext.EXPECT().Argument(0).Return("subdir/DemoFactory").Once()
	mockContext.EXPECT().OptionBool("force").Return(false).Once()
	mockContext.EXPECT().Success("Factory created successfully").Once()
	assert.NoError(t, factoryMakeCommand.Handle(mockContext))
	assert.True(t, file.Exists("database/factories/subdir/demo_factory.go"))
	assert.True(t, file.Contain("database/factories/subdir/demo_factory.go", "package subdir"))
	assert.True(t, file.Contain("database/factories/subdir/demo_factory.go", "type DemoFactory struct"))
	assert.NoError(t, file.Remove("database"))
}
