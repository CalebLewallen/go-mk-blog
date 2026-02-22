---
Page Name: Querying JSON Data from SQL Tables
Published: false
Published Date: 
Author: Caleb Lewallen
RobotsAllowed: true
---

![Feature Image](/static/images/photo-1610986602538-431d65df4385)

If you've worked with data for long enough, you're going to run into something that leaves you wondering , "but, why?". For me, that thing was storing JSON data in a relational table.

Yeah, yeah, I get it, there's a time and place for it, one of you is gonna tell me how using JSONB in Postgres totally changed your life. While that's totally legitimate, sometimes in a role as a data analyst or data engineer, it's difficult to decompose information stored in that format (unless your name is Devin).

For the rest of us, we need a way to be able to query, and sometimes perform complex queries, on large, deeply nested JSON data. What I'm going to show you today, is a way to decompose a JSON object, stored in a relational table, and query it in a tabular format.

# Python to the Rescue!

We're going to use every data professional's best friend, Python, to create our SQL query.

```Python
# Adjusting the function to avoid duplicate yields for primitive values

def get_paths_from_json_object(json_object, current_path='') -> Generator:
    if isinstance(json_object, dict):
        for key, value in json_object.items():
            new_path = f"{current_path}.{key}" if current_path else key
            yield new_path
            yield from get_paths_from_json_object(value, new_path)
    elif isinstance(json_object, list):
        for idx, value in enumerate(json_object):
            new_path = f"{current_path}[{idx}]"
            yield new_path
            yield from get_paths_from_json_object(value, new_path)

# Testing the adjusted function with the same JSON structure
paths_fixed = get_paths_from_json(test_json)
paths_fixed
```
