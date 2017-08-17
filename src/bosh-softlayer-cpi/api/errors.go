package api

import (
	"fmt"
)

type CloudError interface {
	error

	Type() string
}

type RetryableError interface {
	error

	CanRetry() bool
}

type NotSupportedError struct{}

func (e NotSupportedError) Type() string  { return "Bosh::Clouds::NotSupported" }
func (e NotSupportedError) Error() string { return "Not supported" }

type VMNotFoundError struct {
	vmID string
}

func NewVMNotFoundError(vmID string) VMNotFoundError {
	return VMNotFoundError{vmID: vmID}
}

func (e VMNotFoundError) Type() string  { return "Bosh::Clouds::VMNotFound" }
func (e VMNotFoundError) Error() string { return fmt.Sprintf("VM '%s' not found", e.vmID) }

type VMCreationFailedError struct {
	reason   string
	canRetry bool
}

func NewVMCreationFailedError(reason string, canRetry bool) VMCreationFailedError {
	return VMCreationFailedError{reason: reason, canRetry: canRetry}
}

func (e VMCreationFailedError) Type() string   { return "Bosh::Clouds::VMCreationFailed" }
func (e VMCreationFailedError) Error() string  { return fmt.Sprintf("VM failed to create: %v", e.reason) }
func (e VMCreationFailedError) CanRetry() bool { return e.canRetry }

type DiskCreationFailedError struct {
	reason   string
	canRetry bool
}

func NewDiskCreationFailedError(reason string, canRetry bool) DiskCreationFailedError {
	return DiskCreationFailedError{reason: reason, canRetry: canRetry}
}

func (e DiskCreationFailedError) Type() string { return "Bosh::Clouds::DiskCreationFailed" }
func (e DiskCreationFailedError) Error() string {
	return fmt.Sprintf("Disk failed to create: %v", e.reason)
}
func (e DiskCreationFailedError) CanRetry() bool { return e.canRetry }

type NoDiskSpaceError struct {
	diskID   string
	canRetry bool
}

func NewNoDiskSpaceError(diskID string, canRetry bool) NoDiskSpaceError {
	return NoDiskSpaceError{diskID: diskID, canRetry: canRetry}
}

func (e NoDiskSpaceError) Type() string   { return "Bosh::Clouds::NoDiskSpace" }
func (e NoDiskSpaceError) Error() string  { return fmt.Sprintf("Disk '%s' has no space", e.diskID) }
func (e NoDiskSpaceError) CanRetry() bool { return e.canRetry }

type DiskNotAttachedError struct {
	vmID     string
	diskID   string
	canRetry bool
}

func NewDiskNotAttachedError(vmID string, diskID string, canRetry bool) DiskNotAttachedError {
	return DiskNotAttachedError{vmID: vmID, diskID: diskID, canRetry: canRetry}
}

func (e DiskNotAttachedError) Type() string { return "Bosh::Clouds::DiskNotAttached" }
func (e DiskNotAttachedError) Error() string {
	return fmt.Sprintf("Disk '%s' not attached to VM '%s'", e.diskID, e.vmID)
}
func (e DiskNotAttachedError) CanRetry() bool { return e.canRetry }

type DiskNotFoundError struct {
	diskID   string
	canRetry bool
}

func NewDiskNotFoundError(diskID string, canRetry bool) DiskNotFoundError {
	return DiskNotFoundError{diskID: diskID, canRetry: canRetry}
}

func (e DiskNotFoundError) Type() string   { return "Bosh::Clouds::DiskNotFound" }
func (e DiskNotFoundError) Error() string  { return fmt.Sprintf("Disk '%s' not found", e.diskID) }
func (e DiskNotFoundError) CanRetry() bool { return e.canRetry }

type StemcellNotFoundError struct {
	stemcellID string
	canRetry   bool
}

func NewStemcellkNotFoundError(stemcellID string, canRetry bool) StemcellNotFoundError {
	return StemcellNotFoundError{stemcellID: stemcellID, canRetry: canRetry}
}

func (e StemcellNotFoundError) Type() string { return "Bosh::Clouds::StemcellNotFound" }
func (e StemcellNotFoundError) Error() string {
	return fmt.Sprintf("Stemcell '%s' not found", e.stemcellID)
}
func (e StemcellNotFoundError) CanRetry() bool { return e.canRetry }

type HostHaveNotAllowedCredentialError struct {
	vmID     string
	canRetry bool
}

func NewHostHaveNotAllowedCredentialError(vmID string) HostHaveNotAllowedCredentialError {
	return HostHaveNotAllowedCredentialError{vmID: vmID}
}

func (e HostHaveNotAllowedCredentialError) Type() string {
	return "Bosh::Clouds::HostHaveNotAllowedCredential"
}
func (e HostHaveNotAllowedCredentialError) Error() string {
	return fmt.Sprintf("VM '%s' have not allowed access credential", e.vmID)
}
func (e HostHaveNotAllowedCredentialError) CanRetry() bool { return e.canRetry }
