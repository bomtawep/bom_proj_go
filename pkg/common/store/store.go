package store

//var store *session.Store
//
//func InitialStore() {
//	storage, err := coherence.New()
//	if err != nil {
//		panic(err)
//	}
//	store = session.New(session.Config{
//		Expiration: time.Duration(120) * time.Second,
//		Storage:    storage,
//	})
//}
//
//func GetStore(c *fiber.Ctx) (*session.Session, error) {
//	// Get session from store
//	sess, err := store.Get(c)
//	if err != nil {
//		panic(err)
//	}
//	return sess, nil
//}
