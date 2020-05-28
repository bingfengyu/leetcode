package main

import "fmt"

func main() {
	// twitter.PostTweet(1, 1)
	// twitter.Follow(2, 1)
	// twitter.Unfollow(1, 2)
	// fmt.Println(twitter.GetNewsFeed(1))
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	twitter.Unfollow(1, 1)
	fmt.Println(twitter.GetNewsFeed(1))
}

//试题地址://https://leetcode-cn.com/problems/design-twitter/
type Users struct {
	Friend map[int]bool //好友
}

type News struct {
	Users int
	New   int
}

type Twitter struct {
	AllUsers map[int]Users
	AllNews  []News
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	var twitter Twitter
	twitter.AllUsers = map[int]Users{}
	twitter.AllNews = []News{}
	return twitter
}

//初始化用户
func (this *Twitter) Init(userId int) {
	if _, ok := this.AllUsers[userId].Friend[userId]; !ok {
		var users Users
		users.Friend = map[int]bool{}
		this.AllUsers[userId] = users
		this.AllUsers[userId].Friend[userId] = true
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Init(userId)
	var news News
	news.Users = userId
	news.New = tweetId
	this.AllNews = append(this.AllNews, news)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	this.Init(userId)
	news := []int{}
	user := this.AllUsers[userId]
	for i := len(this.AllNews) - 1; i >= 0; i-- {
		if len(news) == 10 {
			break
		}

		if user.Friend[this.AllNews[i].Users] {
			news = append(news, this.AllNews[i].New)
		}
	}
	return news
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.Init(followerId)
	this.AllUsers[followerId].Friend[followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	this.Init(followerId)
	this.AllUsers[followerId].Friend[followeeId] = false
}
