package services

//type CheckAccessService struct {
//	repo repositories.ICheckAccessRepo
//}
//
//func NewCheckAccessService(repo repositories.ICheckAccessRepo) *CheckAccessService {
//	return &CheckAccessService{
//		repo: repo,
//	}
//}
//
//func (s *CheckAccessService) CheckAccessUser(request models.CheckUserAccessRequest, response chan models.CheckUserAccessResponse) {
//	result := s.repo.CheckAccessUser(request)
//
//	response <- models.CheckUserAccessResponse{
//		Code:   result.Code,
//		Status: result.Status,
//		Error:  result.Error,
//	}
//}
