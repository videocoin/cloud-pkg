package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	MinerUserTokenDev = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl90eXBlIjowLCJleHAiOjE1ODk3NjExNTgsInN1YiI6IjZiN2E5Zjk1LWI2ZTYtNDYwOS01NWYyLWRhNGFlZjM2MTg0ZSJ9.feij69LNC72Ernksu98cAWu_8ArXG2kb7-8kgubdMGE"
	MinerUserIDDev    = "6b7a9f95-b6e6-4609-55f2-da4aef36184e"

	MinerUserTokenStage = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl90eXBlIjowLCJleHAiOjE1ODk3MjYzODUsInN1YiI6IjllNDNiMzU4LTRkYzctNDA3ZS01ODY4LWExNGYwNjk0N2ZiZSJ9.9K1Z9WMPl1pQlch1jN9Rmpoo0VKQ6OJkJvRZLLv4VZo"
	MinerUserIDStage    = "9e43b358-4dc7-407e-5868-a14f06947fbe"
)

func TestCreateKeyOnDev(t *testing.T) {
	cli, err := NewClient("https://iam.dev.videocoinapis.com")
	require.NoError(t, err)

	sa, err := cli.CreateServiceAccount(MinerUserTokenDev, MinerUserIDDev)
	require.NoError(t, err, "failed to create service account")

	assert.Equal(t, "service_account", sa.Type)
	assert.Equal(t, MinerUserIDDev, sa.ClientID)
	assert.NotEmpty(t, sa.PrivateKeyID)
	assert.NotEmpty(t, sa.PrivateKey)
}

func TestCreateKeyOnStage(t *testing.T) {
	cli, err := NewClient("https://iam.staging.videocoinapis.com")
	require.NoError(t, err)

	sa, err := cli.CreateServiceAccount(MinerUserTokenStage, MinerUserIDStage)
	require.NoError(t, err, "failed to create service account")

	assert.Equal(t, "service_account", sa.Type)
	assert.Equal(t, MinerUserIDStage, sa.ClientID)
	assert.NotEmpty(t, sa.PrivateKeyID)
	assert.NotEmpty(t, sa.PrivateKey)
}
