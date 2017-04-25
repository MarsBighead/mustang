db.trans.aggregate([
    {$match: {
        "response.request.slot.site.id":"YcsNybZR", 
        "response.created_at":{$gt: ISODate("2017-04-18T16:00:00.441Z")},
        "impression.created_at": {$exists:1},
    }},
    {$group: {
        _id: {
                y: {$year: "$impression.created_at"},
                m: {$month: "$impression.created_at"},
                d: {$dayOfMonth: "$impression.created_at"},
        },
        impression: { $sum: "$stats.impression" }
    }}
]).pretty()
