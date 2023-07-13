# Go-CRUD-API-example

### Preparation

1. Login to https://www.wolframalpha.com/ and developer portal ![setup-1.png](./docs/images/setup-1.png)
2. Create AppID ![setup-2.png](./docs/images/setup-2.png) ![setup-3.png](./docs/images/setup-3.png) and wait 24 hours
3. Save AppID ![setup-10.png](./docs/images/setup-10.png) and use it in WOLFRAM_APP_ID
4. Login to https://wit.ai/
5. Create new app ![setup-4.png](./docs/images/setup-4.png) ![setup-5.png](./docs/images/setup-5.png)
6. Create new intent for app ![setup-6.png](./docs/images/setup-6.png) ![setup-7.png](./docs/images/setup-7.png)
7. Create new entity for app ![setup-8.png](./docs/images/setup-8.png) ![setup-9.png](./docs/images/setup-9.png)
8. Save server access token ![setup-11.png](./docs/images/setup-11.png) and use it in WIT_AI_TOKEN
9. Create an account in Slack
10. Go to https://api.slack.com/apps?new_app=1 and create slack app ![setup-12.png](./docs/images/setup-1.png) ![setup-13.png](./docs/images/setup-2.png) ![setup-3.png](./docs/images/setup-3.png)
11. Get Slack application token and use it in SLACK_APP_TOKEN
12. Enable sockets ![setup-14.png](./docs/images/setup-4.png) ![setup-15.png](./docs/images/setup-5.png) and save token ![setup-16.png](./docs/images/setup-16.png)
13. Enable event subscriptions ![setup-17.png](./docs/images/setup-17.png)
14. Configure Slack Event Subscriptions ![setup-18.png](./docs/images/setup-18.png)
15. Configure Slack OAuth and Permissions with adding scopes ![setup-19.png](./docs/images/setup-19.png) ![setup-20.png](./docs/images/setup-20.png) ![setup-21.png](./docs/images/setup-21.png)
16. Get Slack Bot User OAuth Token and set SLACK_BOT_TOKEN in config file ![setup-22.png](./docs/images/setup-22.png)
17. Add ai-bot to the channel in Slack ![setup-14.png](./docs/images/setup-14.png) ![setup-25.png](./docs/images/setup-26.png)

### Run

```
go run main.go
```

Go to Slack channel where ai-bot was added and ask question with next structure 
```
@wit-ai query for bot - {{TEST_OF_QUESTION}}
```
![result.png](./docs/images/result.png)