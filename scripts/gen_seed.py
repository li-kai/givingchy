import os
import csv
import string
import random
from hashlib import md5
from faker import Faker

"""
Generates the required csv files for seeding the database.
Related files: db/z_seed.sql

To use, simply run: python3 gen_seed.py
"""

dirname = os.path.dirname(__file__)
generator = Faker()

MULTIPLIER = 100
NO_OF_USERS = 10 * MULTIPLIER
NO_OF_PROJECTS = 5 * MULTIPLIER
NO_OF_COMMENTS = 100 * MULTIPLIER
MAX_TAGS_PER_PROJ = 5
MAX_REQUIRED = 10000
CATEGORIES = [
    "Art",
    "Comics",
    "Crafts",
    "Dance",
    "Design",
    "Fashion",
    "Film & Video",
    "Food",
    "Games",
    "Journalism",
    "Music",
    "Photography",
    "Publishing",
    "Technology",
    "Theater",
]

BASE_ALPH = tuple(string.ascii_letters + '_+')
BASE_DICT = dict((c, v) for v, c in enumerate(BASE_ALPH))
BASE_LEN = len(BASE_ALPH)


def base_encode(num):
    if not num:
        return BASE_ALPH[0]

    encoding = ""
    while num:
        num, rem = divmod(num, BASE_LEN)
        encoding = BASE_ALPH[rem] + encoding
    return encoding


print("write_user_csv")
user_csv_path = os.path.join(dirname, '../db/users.csv')
with open(user_csv_path, 'w') as user_csv:
    user_writer = csv.writer(user_csv)
    user_writer.writerow([
        "user_id",
        "email",
        "password",
        "username",
        "total_donation",
        "image",
        "is_admin",
    ])
    user_writer.writerow(["1", "admin", "password", "admin",
                          0, "https://www.gravatar.com/avatar/", "t"])
    for i in range(2, NO_OF_USERS + 1):
        email = generator.ascii_email()
        password = generator.color_name().lower()
        while len(password) < 8:
            password += password
        user_writer.writerow([
            i,
            base_encode(i) + email,
            password,
            generator.user_name(),
            0,
            "https://www.gravatar.com/avatar/" +
            md5(email.encode('utf-8')).hexdigest(),
            "f"
        ])
print("write_user_csv end")

print("write_category_csv")
category_csv_path = os.path.join(dirname, '../db/categories.csv')
with open(category_csv_path, 'w') as category_csv:
    category_writer = csv.writer(category_csv)
    category_writer.writerow(["name", "proj_num"])
    for cate in CATEGORIES:
        category_writer.writerow([cate, 0])
print("write_category_csv end")

print("write_project_csv and write_payment_csv")
project_csv_path = os.path.join(dirname, '../db/projects.csv')
payment_csv_path = os.path.join(dirname, '../db/payments.csv')
with open(project_csv_path, 'w') as project_csv, open(payment_csv_path, 'w') as payment_csv:
    project_writer = csv.writer(project_csv)
    payment_writer = csv.writer(payment_csv)
    project_writer.writerow([
        "project_id",
        "title",
        "user_id",
        "category",
        "description",
        "verified",
        "image",
        "amount_raised",
        "amount_required",
        "start_time",
        "end_time",
    ])
    payment_writer.writerow(["id", "user_id", "project_id", "moment", "amount"])
    payment_id = 1
    for i in range(1, NO_OF_PROJECTS + 1):
        amount_required = random.randint(1, MAX_REQUIRED * 100)
        start_date = generator.past_datetime()
        end_date = generator.date_time_between(start_date, "+30d")
        project_writer.writerow([
            i,
            generator.company(),
            random.randint(2, NO_OF_USERS),
            random.choice(CATEGORIES),
            generator.catch_phrase(),
            "t",
            "https://picsum.photos/320/200/?random" + str(i),
            0,
            "{0:.2f}".format(amount_required / 100),
            start_date,
            end_date,
        ])
        # write payments, nested because payment needs to be between
        # valid period, and amount must tally raised
        left_to_raise = random.randint(0, amount_required)
        while left_to_raise > 0:
            amount = random.randint(1, left_to_raise)
            payment_writer.writerow([
                payment_id,
                random.randint(2, NO_OF_USERS),
                i,
                generator.date_time_between(start_date, end_date),
                "{0:.2f}".format(amount / 100)
            ])
            left_to_raise -= amount
            payment_id += 1
print("write_project_csv and write_payment_csv end")

print("write_tag_csv")
tag_csv_path = os.path.join(dirname, '../db/tags.csv')
with open(tag_csv_path, 'w') as tag_csv:
    tag_writer = csv.writer(tag_csv)
    tag_writer.writerow(
        ["project_id", "tag_name"])
    for i in range(1, NO_OF_PROJECTS + 1):
        countries = set()
        for j in range(random.randint(0, MAX_TAGS_PER_PROJ)):
            countries.add(generator.country())
        for country in countries:
            tag_writer.writerow([
                i,
                country,
            ])
print("write_tag_csv end")

print("write_comment_csv")
comment_csv_path = os.path.join(dirname, '../db/comments.csv')
with open(comment_csv_path, 'w') as comment_csv:
    comment_writer = csv.writer(comment_csv)
    comment_writer.writerow(
        ["id", "user_id", "project_id", "moment", "content"])
    for i in range(1, NO_OF_COMMENTS + 1):
        comment_writer.writerow([
            i,
            random.randint(2, NO_OF_USERS),
            random.randint(1, NO_OF_PROJECTS),
            generator.past_datetime(),
            generator.sentence(),
        ])
print("write_comment_csv end")
