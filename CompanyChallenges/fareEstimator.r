times <- c()
distances <- c()
final <- c()
last <- c()
fareEstimator <- function(ride_time, ride_distance, cost_per_minute, cost_per_mile) {
    size <- length(cost_per_mile)
    times[1:size] <- c(ride_time)
    distances[1:size] <- c(ride_distance)
    d <- data.frame("cpmin" = unlist(cost_per_minute), "time" = times, "cpmile" = unlist(cost_per_mile),"distance" = distances)
    last <- apply(d,1,calc)
    return(last)
    
}

calc <- function(x) {
    final <- c(final, (x[1] * x[2]) + (x[3] * x[4]))
}

