const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[magazzino] magazzini collection - #docs: ", db.mag_magazzini.countDocuments())
print("[magazzino] boxes collection - #docs: ", db.mag_boxes.countDocuments())

