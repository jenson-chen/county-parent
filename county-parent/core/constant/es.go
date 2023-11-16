package constant

const (
	UserBaseIndex   = "user_info"
	CountyBaseIndex = "county_info"

	MappingsOfCounty = "{\n    \"settings\":{\n        \"index\":{\n            \"number_of_shards\":5,\n            \"number_of_replicas\":0\n        }\n    },\n    \"mappings\" : {\n        \"properties\":{\n            \"CountyName\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_smart\"\n            },\n            \"PersonName\":{\n                \"type\":\"keyword\"\n            },\n            \"SchoolCount\":{\n                \"type\":\"integer\"\n            },\n            \"TeacherCount\":{\n                \"type\":\"integer\"\n            },\n\t\t\t\"ClassCount\":{\n                \"type\":\"integer\"\n            },\n            \"EduInvest\":{\n                \"type\":\"float\"\n            },\n\t\t\t\"ItlgInvest\":{\n                \"type\":\"float\"\n            },\n\t\t\t\"AnnualKeyword\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_max_word\"\n            },\n            \"EduKeyword\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_max_word\"\n            },\n\t\t\t\"CtlFunc\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_max_word\"\n            },\n\t\t\t\"NetOperator\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_smart\"\n            },\n\t\t\t\"OnlineCtl\":{\n                \"type\":\"text\",\n                \"analyzer\":\"ik_max_word\"\n            }\n        }\n    }\n}"
)
