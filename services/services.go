package services

import "class-review-backend/repositories"

type Services struct {
    ReviewService IReviewService
    CourseService ICourseService
}

func DefaultServices(repos *repositories.Repositories) *Services {
    return &Services{
        ReviewService: DefaultReviewServices(repos),
        CourseService: DefaultCourseServices(repos),
    }

}
