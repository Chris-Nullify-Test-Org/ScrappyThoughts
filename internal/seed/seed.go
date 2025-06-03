package seed

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

// SeedData initializes the database with dummy data
func SeedData(db *gorm.DB) error {
	// Create users with different roles
	users := []models.User{
		{
			Username: "sarah_admin",
			Email:    "sarah.johnson@scrappythoughts.com",
			Password: "Admin2023!",
			Role:     "admin",
		},
		{
			Username: "mike_mod",
			Email:    "mike.wilson@scrappythoughts.com",
			Password: "Mod2023!",
			Role:     "moderator",
		},
		{
			Username: "lisa_mod",
			Email:    "lisa.garcia@scrappythoughts.com",
			Password: "Mod2023!",
			Role:     "moderator",
		},
		{
			Username: "tech_writer",
			Email:    "alex.chen@gmail.com",
			Password: "Password123",
			Role:     "user",
		},
		{
			Username: "travel_bug",
			Email:    "emma.taylor@outlook.com",
			Password: "Password123",
			Role:     "user",
		},
		{
			Username: "foodie_forever",
			Email:    "james.smith@yahoo.com",
			Password: "Password123",
			Role:     "user",
		},
		{
			Username: "bookworm",
			Email:    "sophia.brown@hotmail.com",
			Password: "Password123",
			Role:     "user",
		},
		{
			Username: "fitness_fanatic",
			Email:    "david.miller@gmail.com",
			Password: "Password123",
			Role:     "user",
			IsBanned: true,
		},
	}

	// Create users
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	// Create posts
	posts := []models.Post{
		{
			Title:    "Welcome to ScrappyThoughts: A Community for Authentic Voices",
			Content:  "We're excited to launch ScrappyThoughts, a platform where writers can share their authentic perspectives without the pressure of perfection. Whether you're a seasoned author or just starting your writing journey, this is your space to be real, raw, and relatable.\n\nOur mission is simple: to create a community where genuine thoughts and experiences are valued over polished prose. We believe that the most impactful stories often come from the heart, not from carefully crafted sentences.\n\nJoin us as we build this community together. Share your thoughts, engage with others, and let's create something special.",
			AuthorID: 1, // sarah_admin
			Status:   "active",
		},
		{
			Title:    "The Rise of Microservices: A Developer's Perspective",
			Content:  "In the ever-evolving landscape of software development, microservices have emerged as a compelling architectural style. As someone who's been working with this approach for the past three years, I wanted to share my experiences and insights.\n\nMicroservices offer several advantages: they allow teams to work independently, enable technology diversity, and provide better fault isolation. However, they also introduce complexity in deployment, monitoring, and inter-service communication.\n\nIn this post, I'll explore the pros and cons of microservices, share some real-world examples, and discuss when this architecture might be the right choice for your project.",
			AuthorID: 4, // tech_writer
			Status:   "active",
		},
		{
			Title:    "Hidden Gems: My Journey Through Southeast Asia",
			Content:  "After six months of backpacking through Southeast Asia, I've discovered that the most memorable experiences often come from the least expected places. While everyone visits the famous temples of Angkor Wat or the beaches of Bali, some of my most cherished memories are from tiny villages and local markets.\n\nIn this post, I'll share five lesser-known destinations that stole my heart, along with practical tips for travelers who want to venture off the beaten path. From a family-run homestay in northern Vietnam to a secret waterfall in Laos, these places offer authentic experiences without the crowds.",
			AuthorID: 5, // travel_bug
			Status:   "active",
		},
		{
			Title:    "The Art of Sourdough: A Beginner's Guide",
			Content:  "There's something magical about creating bread from just flour, water, and salt. After two years of experimenting with sourdough, I've finally developed a reliable method that even beginners can follow.\n\nIn this comprehensive guide, I'll walk you through the entire process: creating and maintaining a starter, understanding the fermentation process, shaping your dough, and achieving that perfect crust and crumb. I'll also share some common mistakes I made along the way so you can avoid them.\n\nWhether you're a complete novice or have been baking for years, this guide will help you create delicious, artisanal sourdough bread in your own kitchen.",
			AuthorID: 6, // foodie_forever
			Status:   "active",
		},
		{
			Title:    "Why I'm Rereading the Classics in 2023",
			Content:  "In an age of endless new releases and digital distractions, I've made a conscious decision to revisit the literary classics this year. There's something profoundly different about reading these works as an adult compared to my first encounters with them in school.\n\nSo far, I've revisited 'To Kill a Mockingbird,' '1984,' and 'The Great Gatsby.' Each reading has revealed new layers of meaning and relevance to our contemporary world. In this post, I'll share my reflections on these timeless works and why I believe they still matter in 2023.",
			AuthorID: 7, // bookworm
			Status:   "active",
		},
		{
			Title:    "From Couch to Marathon: My Year-Long Journey",
			Content:  "Last January, I couldn't run for more than five minutes without getting winded. This past weekend, I crossed the finish line of my first marathon. The transformation has been both physical and mental.\n\nIn this post, I'll share my training approach, the challenges I faced, and the lessons I learned along the way. I'll also provide a beginner-friendly training plan that helped me build endurance gradually and avoid injury.\n\nWhether you're considering your first 5K or dreaming of a marathon, I hope my experience inspires you to take that first step toward your running goals.",
			AuthorID: 8, // fitness_fanatic
			Status:   "active",
		},
		{
			Title:    "The Ethics of Artificial Intelligence: A Critical Discussion",
			Content:  "As AI systems become increasingly sophisticated and integrated into our daily lives, we face important ethical questions that demand our attention. From algorithmic bias to job displacement, the implications of AI development extend far beyond technical considerations.\n\nIn this post, I'll examine some of the most pressing ethical concerns surrounding AI, including privacy, transparency, accountability, and the potential for misuse. I'll also discuss frameworks for responsible AI development and the role of regulation in ensuring these technologies benefit society as a whole.\n\nThis isn't just a theoretical discussion—it's a call to action for developers, policymakers, and citizens to engage with these issues before they're decided for us.",
			AuthorID: 4, // tech_writer
			Status:   "active",
		},
		{
			Title:    "Sustainable Travel: How to See the World Without Destroying It",
			Content:  "Travel has the power to broaden our horizons and connect us with different cultures, but it also has a significant environmental impact. After witnessing the effects of overtourism in several destinations, I've become passionate about sustainable travel practices.\n\nIn this post, I'll share practical tips for reducing your travel footprint, from choosing eco-friendly accommodations to supporting local communities. I'll also discuss the concept of 'slow travel' and how taking a more mindful approach to exploration can lead to more meaningful experiences.\n\nTravel doesn't have to be a choice between seeing the world and protecting it. With some planning and consideration, we can do both.",
			AuthorID: 5, // travel_bug
			Status:   "active",
		},
		{
			Title:    "The Hidden Costs of Fast Fashion",
			Content:  "I recently conducted a month-long experiment: I tracked every piece of clothing I purchased, its cost, how many times I wore it, and how long it lasted. The results were eye-opening and led me to completely rethink my approach to fashion.\n\nIn this post, I'll share my findings about the true cost of fast fashion—not just in dollars, but in environmental impact, labor conditions, and personal satisfaction. I'll also provide practical advice for building a more sustainable wardrobe that's both stylish and ethical.\n\nIt's time to move beyond the quick fixes of fast fashion and toward a more thoughtful approach to what we wear.",
			AuthorID: 7, // bookworm
			Status:   "active",
		},
		{
			Title:    "Controversial Opinion: Why I Believe in Moderate Exercise",
			Content:  "In a fitness world dominated by extreme workouts and 'no pain, no gain' mentality, I'm here to make the case for moderation. After years of pushing my body to its limits, I've discovered that a balanced approach to exercise yields better long-term results.\n\nIn this post, I'll share my personal journey from overtraining to a more sustainable fitness routine. I'll discuss the science behind moderate exercise, its benefits for longevity, and how to find the right balance for your body and goals.\n\nThis isn't about being lazy—it's about being smart about how we train and recognizing that more isn't always better.",
			AuthorID: 8, // fitness_fanatic
			Status:   "active",
		},
		{
			Title:    "The Psychology of Food: Why We Crave What We Crave",
			Content:  "Have you ever wondered why you can't resist that chocolate bar after a stressful day, or why certain foods trigger powerful memories? As a chef and food psychology enthusiast, I've spent years exploring the fascinating connection between our minds and our plates.\n\nIn this post, I'll dive into the science behind food cravings, emotional eating, and how our early experiences shape our food preferences. I'll also share practical strategies for developing a healthier relationship with food.\n\nUnderstanding the psychology of eating can help us make more mindful choices and enjoy our meals on a deeper level.",
			AuthorID: 6, // foodie_forever
			Status:   "active",
		},
		{
			Title:    "The Future of Remote Work: Lessons from Two Years of WFH",
			Content:  "After two years of working remotely, I've experienced both the incredible benefits and the unexpected challenges of this new work paradigm. As someone who manages a distributed team, I've had a front-row seat to how remote work is reshaping our professional lives.\n\nIn this post, I'll share insights about productivity, collaboration, and maintaining work-life boundaries in a remote setting. I'll also discuss the technologies and practices that have been most effective for my team.\n\nWhether you're a remote work veteran or just starting out, I hope these lessons will help you navigate this evolving landscape more effectively.",
			AuthorID: 4, // tech_writer
			Status:   "active",
		},
		{
			Title:    "The Art of Mindful Reading in a Distracted World",
			Content:  "In an age of endless notifications and infinite scrolling, the simple act of reading a book has become a challenge. As someone who reads for both pleasure and professional development, I've developed strategies for maintaining focus and getting more from my reading time.\n\nIn this post, I'll share my approach to mindful reading, including techniques for creating the right environment, managing digital distractions, and engaging more deeply with the text. I'll also recommend some books that have particularly benefited from this approach.\n\nReading doesn't have to be another casualty of our distracted age. With intention and practice, we can reclaim this valuable activity.",
			AuthorID: 7, // bookworm
			Status:   "active",
		},
		{
			Title:    "The Hidden Dangers of Extreme Dieting",
			Content:  "I recently came across a post promoting an extremely restrictive diet plan that promised rapid weight loss. As someone who's recovered from disordered eating, I felt compelled to address the dangerous misinformation being spread.\n\nIn this post, I'll discuss the physical and psychological risks of extreme dieting, including metabolic damage, nutrient deficiencies, and the development of eating disorders. I'll also share evidence-based approaches to healthy weight management.\n\nIt's time to challenge the harmful narratives around dieting and promote a more balanced, sustainable approach to nutrition and body image.",
			AuthorID: 8, // fitness_fanatic
			Status:   "hidden",
		},
		{
			Title:    "Why I'm Taking a Break from Social Media",
			Content:  "After five years of being constantly connected, I've decided to take a three-month break from all social media platforms. This decision wasn't made lightly, but after noticing significant changes in my attention span, anxiety levels, and overall well-being, I knew something had to change.\n\nIn this post, I'll share my observations about how social media has affected my mental health, relationships, and productivity. I'll also discuss the specific steps I'm taking to maintain connections without the platforms and how I'm filling the time I previously spent scrolling.\n\nThis isn't a permanent goodbye, but rather an experiment in digital minimalism and a quest to find a healthier balance with technology.",
			AuthorID: 5, // travel_bug
			Status:   "active",
		},
		{
			Title:    "The Ethics of Travel Photography",
			Content:  "As a travel photographer, I've grappled with questions about representation, consent, and the impact of my work on the communities I document. These aren't just theoretical concerns—they affect real people and places.\n\nIn this post, I'll explore the ethical considerations of travel photography, from how we portray different cultures to the environmental impact of our presence in fragile ecosystems. I'll also share my evolving approach to capturing images that respect both subjects and viewers.\n\nPhotography has the power to shape perceptions and influence behavior. With that power comes responsibility—to be thoughtful about what and how we photograph.",
			AuthorID: 5, // travel_bug
			Status:   "active",
		},
		{
			Title:    "The Science of Flavor: Understanding Taste and Aroma",
			Content:  "Have you ever wondered why certain flavor combinations work so well together, or why some people love cilantro while others find it soapy? As a chef and food science enthusiast, I'm fascinated by the complex interplay of taste, smell, and texture that creates our experience of flavor.\n\nIn this post, I'll break down the science behind our sense of taste, including the five basic tastes, how aroma enhances flavor, and the role of texture in our enjoyment of food. I'll also share some surprising facts about individual differences in taste perception.\n\nUnderstanding the science of flavor can transform how we cook, eat, and appreciate food. It's not just about following recipes—it's about understanding the principles that make them work.",
			AuthorID: 6, // foodie_forever
			Status:   "active",
		},
	}

	// Create posts
	for _, post := range posts {
		if err := db.Create(&post).Error; err != nil {
			return err
		}
	}

	// Create comments
	comments := []models.Comment{
		{
			Content:  "This is exactly what I've been looking for! A platform that values authenticity over perfection. Looking forward to being part of this community.",
			PostID:   1,
			AuthorID: 4, // tech_writer
		},
		{
			Content:  "Thank you for creating this space. I've always felt intimidated by other blogging platforms that seem to prioritize polished content over genuine thoughts.",
			PostID:   1,
			AuthorID: 7, // bookworm
		},
		{
			Content:  "Your microservices article really helped clarify some concepts I've been struggling with. The real-world examples were particularly valuable.",
			PostID:   2,
			AuthorID: 5, // travel_bug
		},
		{
			Content:  "I've been considering microservices for my startup, but was concerned about the complexity. Your article provided a balanced perspective that helped me make a more informed decision.",
			PostID:   2,
			AuthorID: 8, // fitness_fanatic
		},
		{
			Content:  "Your Southeast Asia post brought back so many memories! I completely agree about the hidden gems. The family-run homestay you mentioned in Vietnam was one of my favorite experiences too.",
			PostID:   3,
			AuthorID: 6, // foodie_forever
		},
		{
			Content:  "I've been dreaming of traveling to Southeast Asia, and your post has given me the confidence to venture off the beaten path. Thank you for the practical tips!",
			PostID:   3,
			AuthorID: 7, // bookworm
		},
		{
			Content:  "Your sourdough guide is the most comprehensive I've found online. I've been struggling with my starter for months, but your instructions have finally helped me succeed!",
			PostID:   4,
			AuthorID: 5, // travel_bug
		},
		{
			Content:  "I've always been intimidated by sourdough, but your beginner-friendly approach makes it seem achievable. Looking forward to trying your method this weekend.",
			PostID:   4,
			AuthorID: 8, // fitness_fanatic
		},
		{
			Content:  "Your post about rereading classics resonates with me so much. I recently revisited 'To Kill a Mockingbird' and was struck by how differently I interpreted it as an adult.",
			PostID:   5,
			AuthorID: 4, // tech_writer
		},
		{
			Content:  "I've been thinking about revisiting some classics, and your post has inspired me to create a reading list. Any recommendations for which classic to start with?",
			PostID:   5,
			AuthorID: 6, // foodie_forever
		},
		{
			Content:  "Your marathon journey is incredibly inspiring! I've been thinking about training for my first half marathon, and your post has given me the confidence to start.",
			PostID:   6,
			AuthorID: 7, // bookworm
		},
		{
			Content:  "I appreciate your balanced approach to fitness. Too many fitness influencers promote extreme workouts that aren't sustainable long-term. Your perspective is refreshing.",
			PostID:   6,
			AuthorID: 5, // travel_bug
		},
		{
			Content:  "The ethical considerations you've raised about AI are crucial. As someone working in tech, I think we need more discussions like this to ensure we're developing these technologies responsibly.",
			PostID:   7,
			AuthorID: 8, // fitness_fanatic
		},
		{
			Content:  "Your post about sustainable travel really made me think about my own travel habits. I've committed to making some changes based on your suggestions. Thank you for the eye-opening perspective.",
			PostID:   8,
			AuthorID: 4, // tech_writer
		},
		{
			Content:  "I've been struggling with emotional eating for years, and your post about food psychology was incredibly insightful. The strategies you've shared are practical and seem achievable.",
			PostID:   11,
			AuthorID: 7, // bookworm
		},
		{
			Content:  "Your post about remote work really captures the challenges and opportunities I've experienced. The section about maintaining boundaries between work and personal life was particularly helpful.",
			PostID:   12,
			AuthorID: 5, // travel_bug
		},
		{
			Content:  "I've been considering a social media break for a while, and your post has convinced me to take the plunge. Thank you for sharing your experience and insights.",
			PostID:   14,
			AuthorID: 6, // foodie_forever
		},
		{
			Content:  "This content promotes harmful dieting practices that could lead to eating disorders. I'm flagging this for review.",
			PostID:   13,
			AuthorID: 2, // mike_mod
		},
	}

	// Create comments
	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			return err
		}
	}

	// Create some likes
	likes := []models.UserLikedPost{
		{
			UserID: 4, // tech_writer
			PostID: 3, // Hidden Gems: My Journey Through Southeast Asia
		},
		{
			UserID: 5, // travel_bug
			PostID: 2, // The Rise of Microservices
		},
		{
			UserID: 6, // foodie_forever
			PostID: 5, // Why I'm Rereading the Classics
		},
		{
			UserID: 7, // bookworm
			PostID: 6, // From Couch to Marathon
		},
		{
			UserID: 8, // fitness_fanatic
			PostID: 4, // The Art of Sourdough
		},
		{
			UserID: 2, // mike_mod
			PostID: 1, // Welcome to ScrappyThoughts
		},
		{
			UserID: 3, // lisa_mod
			PostID: 1, // Welcome to ScrappyThoughts
		},
		{
			UserID: 4, // tech_writer
			PostID: 8, // Sustainable Travel
		},
		{
			UserID: 5, // travel_bug
			PostID: 7, // The Ethics of AI
		},
		{
			UserID: 6,  // foodie_forever
			PostID: 11, // The Psychology of Food
		},
		{
			UserID: 7,  // bookworm
			PostID: 12, // The Art of Mindful Reading
		},
		{
			UserID: 8,  // fitness_fanatic
			PostID: 13, // The Hidden Dangers of Extreme Dieting
		},
	}

	// Create likes
	for _, like := range likes {
		if err := db.Create(&like).Error; err != nil {
			return err
		}
	}

	// Update like counts
	for _, like := range likes {
		if err := db.Model(&models.Post{}).Where("id = ?", like.PostID).
			UpdateColumn("likes_count", gorm.Expr("likes_count + ?", 1)).Error; err != nil {
			return err
		}
	}

	fmt.Println("\n=== JWT Tokens for Seeded Users ===")
	var allUsers []models.User
	if err := db.Find(&allUsers).Error; err != nil {
		return err
	}

	for _, user := range allUsers {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
		})

		tokenString, err := token.SignedString([]byte("your-secret-key")) // In production, use environment variable
		if err != nil {
			return fmt.Errorf("error generating token for user %s: %w", user.Username, err)
		}

		fmt.Printf("User: %s (ID: %d, Role: %s)\nToken: %s\n\n", user.Username, user.ID, user.Role, tokenString)
	}
	fmt.Println("===================================")

	return nil
}
