curl -s https://ytrack.learn.ynov.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jtom\"}}){id}}"}'|jq '.[].user[0].id'