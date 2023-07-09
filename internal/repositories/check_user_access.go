package repositories

//type CheckAccessRepo struct {
//	db *db.DB
//}
//
//func NewCheckAccessRepo(db *db.DB) *CheckAccessRepo {
//	return &CheckAccessRepo{
//		db: db,
//	}
//}
//
//func (r *CheckAccessRepo) CheckAccessUser(request models.CheckUserAccessRequest) models.CheckUserAccessResponse {
//	exists, err := r.db.CheckUserAccess(
//		request.UserID, request.OKUD, request.ReportID, request.ChapterID, request.AllowedOP)
//	if err != nil {
//		e := strings.Split(err.Error(), ": ")
//		src, msg := e[0], e[1]
//		return models.CheckUserAccessResponse{
//			Code:   400,
//			Status: checkAccessUserFail,
//			Error: &models.Err{
//				Source:  src,
//				Message: msg,
//			},
//		}
//	}
//
//	if !exists {
//		return models.CheckUserAccessResponse{
//			Code:   404,
//			Status: userNotFoundMsg,
//			Error: &models.Err{
//				Source:  "pq",
//				Message: userNotFoundMsg,
//			},
//		}
//	}
//
//	return models.CheckUserAccessResponse{
//		Code:   200,
//		Status: checkAccessUserSuccess,
//		Error:  nil,
//	}
//}
