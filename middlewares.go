package main

func middlewareLoggedIn(handler CommandHandlerWithUser) CommandHandler {
	return func(s *State, cmd Command) error {
		user, err := s.userService.FetchUserByName(s.config.CURRENT_USER)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}