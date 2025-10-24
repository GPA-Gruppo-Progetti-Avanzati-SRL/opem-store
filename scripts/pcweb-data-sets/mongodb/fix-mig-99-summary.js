print(new Date() + " #################### - fix-mig-99-summary.js")
db.opem_system.aggregate([
    {
        $group:
            {
                _id: "$_et",
                num_records: { $sum: 1 }
            }
    }
]);