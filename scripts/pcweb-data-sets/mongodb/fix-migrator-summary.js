db.opem_system.aggregate([
    {
        $group:
            {
                _id: "$_et",
                num_records: { $sum: 1 }
            }
    }
]);