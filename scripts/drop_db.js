const AWS = require('aws-sdk');

AWS.config.update({
    region: "sa-east-1",
});

const dynamoDB = new AWS.DynamoDB.DocumentClient();
const tableName = 'luna-db';

const scanParams = {
    TableName: tableName,
};

const deleteItems = async (items) => {
    for (const item of items) {
        const deleteParams = {
            TableName: tableName,
            Key: {
                pk: item.pk,
                sk: item.sk,
            },
        };

        try {
            await dynamoDB.delete(deleteParams).promise();
            console.log(`Deleted item with pk: ${item.pk}`);
        } catch (error) {
            console.error(`Could not delete item with pk: ${item.pk}. Error: ${error}`);
        }
    }
};

dynamoDB.scan(scanParams, (error, data) => {
    if (error) {
        console.error("Unable to scan the table. Error:", JSON.stringify(error, null, 2));
    } else {
        console.log("Scan succeeded.");
        deleteItems(data.Items);
    }
});