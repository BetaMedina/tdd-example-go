conn = new Mongo();
db = conn.getDB("tdd");


db.statement.createIndex({ "client.id": 1 }, { unique: false });
db.member.createIndex({ "id": 1 }, { unique: true });

db.member.insert({
  "id": 1,
  "limit": 100000,
  "amount": 0
});

