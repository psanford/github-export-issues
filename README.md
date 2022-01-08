# github-export-issues

This is a cli tool to export issues and comments to a json stream.

It might be useful for archival purposes.

## Usage

If you are exporting from a large repo you'll probably want to provide a personal access tokens to avoid github's low rate limits.

```
./github-export-issues --auth_token ghp_XXXXXXXXXXXXXXXXXXXXXXX --pretty https://github.com/psanford/wormhole-william

{
  "issue": {
    "id": 1090660350,
    "number": 67,
    "state": "closed",
    "locked": false,
    "title": "Add new gio File explorer, so that a user can pick a file on any OS",
    "body": "https://github.com/gioui/gio-x/pull/6\r\n\r\nNot quite done yet.\r\n\r\n",
    "author_association": "NONE",
    "user": {
      "login": "gedw99",
      "id": 53147028,
      "node_id": "MDQ6VXNlcjUzMTQ3MDI4",
      "avatar_url": "https://avatars.githubusercontent.com/u/53147028?v=4",
      "html_url": "https://github.com/gedw99",
      "gravatar_id": "",
      "type": "User",
      "site_admin": false,
      "url": "https://api.github.com/users/gedw99",
      "events_url": "https://api.github.com/users/gedw99/events{/privacy}",
      "following_url": "https://api.github.com/users/gedw99/following{/other_user}",
      "followers_url": "https://api.github.com/users/gedw99/followers",
      "gists_url": "https://api.github.com/users/gedw99/gists{/gist_id}",
      "organizations_url": "https://api.github.com/users/gedw99/orgs",
      "received_events_url": "https://api.github.com/users/gedw99/received_events",
      "repos_url": "https://api.github.com/users/gedw99/repos",
      "starred_url": "https://api.github.com/users/gedw99/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/gedw99/subscriptions"
    },
    "comments": 1,
    "closed_at": "2021-12-29T17:44:45Z",
    "created_at": "2021-12-29T17:35:33Z",
    "updated_at": "2021-12-29T17:44:46Z",
    "url": "https://api.github.com/repos/psanford/wormhole-william/issues/67",
    "html_url": "https://github.com/psanford/wormhole-william/issues/67",
    "comments_url": "https://api.github.com/repos/psanford/wormhole-william/issues/67/comments",
    "events_url": "https://api.github.com/repos/psanford/wormhole-william/issues/67/events",
    "labels_url": "https://api.github.com/repos/psanford/wormhole-william/issues/67/labels{/name}",
    "repository_url": "https://api.github.com/repos/psanford/wormhole-william",
    "reactions": {
      "total_count": 0,
      "+1": 0,
      "-1": 0,
      "laugh": 0,
      "confused": 0,
      "heart": 0,
      "hooray": 0,
      "rocket": 0,
      "eyes": 0,
      "url": "https://api.github.com/repos/psanford/wormhole-william/issues/67/reactions"
    },
    "node_id": "I_kwDODDe8_c5BAif-"
  },
  "comments": [
    {
      "id": 1002706949,
      "node_id": "IC_kwDODDe8_c47xBgF",
      "body": "Maybe you meant to open this on the wormhole-william-mobile project? We already have https://github.com/psanford/wormhole-william-mobile/issues/14",
      "user": {
        "login": "psanford",
        "id": 33375,
        "node_id": "MDQ6VXNlcjMzMzc1",
        "avatar_url": "https://avatars.githubusercontent.com/u/33375?v=4",
        "html_url": "https://github.com/psanford",
        "gravatar_id": "",
        "type": "User",
        "site_admin": false,
        "url": "https://api.github.com/users/psanford",
        "events_url": "https://api.github.com/users/psanford/events{/privacy}",
        "following_url": "https://api.github.com/users/psanford/following{/other_user}",
        "followers_url": "https://api.github.com/users/psanford/followers",
        "gists_url": "https://api.github.com/users/psanford/gists{/gist_id}",
        "organizations_url": "https://api.github.com/users/psanford/orgs",
        "received_events_url": "https://api.github.com/users/psanford/received_events",
        "repos_url": "https://api.github.com/users/psanford/repos",
        "starred_url": "https://api.github.com/users/psanford/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/psanford/subscriptions"
      },
      "reactions": {
        "total_count": 0,
        "+1": 0,
        "-1": 0,
        "laugh": 0,
        "confused": 0,
        "heart": 0,
        "hooray": 0,
        "rocket": 0,
        "eyes": 0,
        "url": "https://api.github.com/repos/psanford/wormhole-william/issues/comments/1002706949/reactions"
      },
      "created_at": "2021-12-29T17:44:45Z",
      "updated_at": "2021-12-29T17:44:45Z",
      "author_association": "OWNER",
      "url": "https://api.github.com/repos/psanford/wormhole-william/issues/comments/1002706949",
      "html_url": "https://github.com/psanford/wormhole-william/issues/67#issuecomment-1002706949",
      "issue_url": "https://api.github.com/repos/psanford/wormhole-william/issues/67"
    }
  ]
}

```
