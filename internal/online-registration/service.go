package online_registration

// Required Implementation Details (Actions)
type OnlineRegistrationService interface {
	// Search Record By id/uuid
	Search(uuid string) (*OnlineRegistration, error)
	// Get All Records
	Index() ([]*OnlineRegistration, error)
	// Save Record
	Store(postData *OnlineRegistration) error
	// Edit/Update Record
	Update(uuuid string) (*OnlineRegistration, error)
	// Delete Record
	Delete(uuid string) (*OnlineRegistration, error)
}