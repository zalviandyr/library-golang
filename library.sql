--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0
-- Dumped by pg_dump version 17.4

-- Started on 2025-11-01 08:52:58 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 221 (class 1259 OID 158215)
-- Name: author_books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.author_books (
    book_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    author_id uuid DEFAULT public.uuid_generate_v4() NOT NULL
);


ALTER TABLE public.author_books OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 158185)
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    biography text NOT NULL
);


ALTER TABLE public.authors OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 158201)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    title text NOT NULL,
    description text NOT NULL,
    isbn text NOT NULL,
    published_at timestamp with time zone NOT NULL,
    publisher_id uuid NOT NULL
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 158193)
-- Name: publishers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.publishers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text NOT NULL,
    address text NOT NULL,
    phone text NOT NULL,
    website text NOT NULL
);


ALTER TABLE public.publishers OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 163043)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    username text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3442 (class 0 OID 158215)
-- Dependencies: 221
-- Data for Name: author_books; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.author_books (book_id, author_id) VALUES ('c767efd6-70cd-4eff-9e51-cab58c8d851e', '3f70f3e2-2a77-45d7-8f8b-d53a14ba6ff7');
INSERT INTO public.author_books (book_id, author_id) VALUES ('c767efd6-70cd-4eff-9e51-cab58c8d851e', '20b1f029-b9f1-4ee2-9b75-2b41d67e80b5');
INSERT INTO public.author_books (book_id, author_id) VALUES ('5932aa91-7852-4602-9397-8de4c2077141', '2588cd72-0012-49c9-b6a6-d8e4e0644e83');
INSERT INTO public.author_books (book_id, author_id) VALUES ('5932aa91-7852-4602-9397-8de4c2077141', '69b7cfe0-89b3-4e25-bad9-c02be25082d4');
INSERT INTO public.author_books (book_id, author_id) VALUES ('c5002f8d-4d05-4550-b640-4351382da269', 'fe5e5818-fdfb-481c-aa71-00469de4764b');
INSERT INTO public.author_books (book_id, author_id) VALUES ('c5002f8d-4d05-4550-b640-4351382da269', '20b1f029-b9f1-4ee2-9b75-2b41d67e80b5');
INSERT INTO public.author_books (book_id, author_id) VALUES ('d82af8d3-b24b-4da0-a3a4-a712fc4b8823', '547c33be-6588-4f31-a43a-4074bf7df29f');
INSERT INTO public.author_books (book_id, author_id) VALUES ('aaf63bb8-4d0f-4550-bd3d-d4fd40d3506d', '5bdad376-22f1-4a8e-8139-fb6fb5b9b840');
INSERT INTO public.author_books (book_id, author_id) VALUES ('aaf63bb8-4d0f-4550-bd3d-d4fd40d3506d', 'f0fe8024-0d51-45f0-af8a-6bc7289f99e2');
INSERT INTO public.author_books (book_id, author_id) VALUES ('1cd15f16-c533-4cc3-ae1f-58dc26e678c8', '0cabfefc-6729-452b-beba-456fa63b555d');


--
-- TOC entry 3439 (class 0 OID 158185)
-- Dependencies: 218
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('20b1f029-b9f1-4ee2-9b75-2b41d67e80b5', '2025-11-01 03:02:16.696541+00', '2025-11-01 03:02:16.696541+00', 'Christophe', 'Mertz', 'Sapiente numquam quo iure modi quia. Iusto aperiam dolores consequatur deleniti vitae. Voluptatem odio qui excepturi quisquam quam. Voluptatem odio praesentium dolor a temporibus. Nesciunt omnis quasi dolorem quo nihil. Magni quo est est error facere. Eius quibusdam tempore voluptatem eum enim. Numquam et iure accusamus minus nihil.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('e659870d-fc23-4a19-9b84-c44a457528c9', '2025-11-01 03:02:16.820303+00', '2025-11-01 03:02:16.820303+00', 'Blanche', 'Schumm', 'Soluta quia distinctio velit error est. Fuga omnis sequi doloribus et quaerat. Accusantium fuga et occaecati sit dicta. Ea corporis aut sed consequatur cupiditate. Quos quia aperiam accusantium sunt quisquam. Eum qui harum voluptatem sapiente fugiat. Ullam asperiores velit rerum quia at.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('b47c29e9-7e14-48af-a1ed-7af1b41ce8eb', '2025-11-01 03:02:16.851614+00', '2025-11-01 03:02:16.851614+00', 'Ned', 'Bernhard', 'Ex delectus ut facere adipisci consequatur. In qui soluta est non nobis. Rem vitae suscipit qui consectetur quibusdam. Qui exercitationem quasi error deserunt magni.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('fe5c7d59-05ed-422e-a65c-11f31910db4e', '2025-11-01 03:02:17.09579+00', '2025-11-01 03:02:17.09579+00', 'Kari', 'Nolan', 'Ut ipsum soluta officiis necessitatibus corporis. Veniam sapiente nam quo id facilis. Recusandae dolor sunt explicabo quis autem. Quas necessitatibus eveniet consectetur quisquam id. Quo voluptatem vel consequuntur nemo ut. Quasi enim libero explicabo in est. Repellat vero ut facere maiores enim.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('fbd67303-f00e-4f57-9901-3a89a6b5a2ad', '2025-11-01 03:06:42.59305+00', '2025-11-01 03:06:42.59305+00', 'Filiberto', 'Stiedemann', 'Cupiditate et adipisci nihil repellendus perferendis. Aut sunt est beatae et nobis. Ab facere ipsum corrupti tempore beatae. Architecto et officia maiores est voluptatem. Vero alias dolorum dolorem corporis sunt. Id aliquam ducimus magnam exercitationem vel. Libero et voluptatum voluptatem a quo. Non sit repudiandae quam recusandae quaerat. Molestiae nobis officia dolores saepe porro. Rerum similique natus id deleniti sed.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('8d359f3f-690d-4b2d-95d8-0e955e10f0fb', '2025-11-01 03:06:42.59305+00', '2025-11-01 03:06:42.59305+00', 'Margarete', 'Kiehn', 'Cum sequi harum magnam in omnis. Facere impedit aliquid qui laboriosam facilis. Tempore saepe quia aut id necessitatibus. Explicabo veniam qui iusto vel hic. Eum facilis neque amet repudiandae qui. Velit corrupti molestias exercitationem ut rerum. Unde corrupti ad debitis laboriosam vitae. Corporis est mollitia dolore reiciendis illo.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('1709080b-634e-493e-b9d5-f5b019b09ed5', '2025-11-01 03:06:42.59305+00', '2025-11-01 03:06:42.59305+00', 'Maurine', 'Reinger', 'Qui quaerat minima corporis et delectus. Placeat voluptatem officiis voluptatem veniam nesciunt. Ipsum delectus id minus quidem molestias. Et corrupti facere explicabo minima ea. Sed error ut repellat modi distinctio. Eum recusandae dolorem et tempora voluptatem.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('84f8e163-4065-4130-8c68-364dd7c01462', '2025-11-01 03:06:42.59305+00', '2025-11-01 03:06:42.59305+00', 'Guido', 'Zemlak', 'Dolorem ut aliquid at mollitia voluptates. Eaque voluptatem architecto officia voluptatum est. Facilis sunt dolor ea praesentium quam. Ad recusandae quidem in veritatis omnis. Et ratione ut molestiae esse debitis. Nihil sit ducimus maxime est molestias. Placeat ipsum voluptatum quas beatae in. Doloremque quam quia voluptas debitis ut.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('41e8e4a5-01bf-4aaf-9775-4e125dc63acf', '2025-11-01 03:06:42.59305+00', '2025-11-01 03:06:42.59305+00', 'Annabelle', 'Simonis', 'Quia natus vel nihil in neque. Aut sit odit nostrum exercitationem consequatur. Illo blanditiis non cum nihil est. Magnam excepturi qui aliquid aut debitis.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('69b7cfe0-89b3-4e25-bad9-c02be25082d4', '2025-11-01 03:12:41.190826+00', '2025-11-01 03:12:41.190826+00', 'Felicita', 'Stoltenberg', 'Adipisci quasi est eum dolores ut. Ipsa esse atque et praesentium sunt. Voluptatem deserunt dolores laboriosam quisquam soluta. Repellat voluptatem eligendi voluptatem consequatur ducimus. Ratione et impedit eos esse rem. Et et molestiae sint necessitatibus voluptas.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('547c33be-6588-4f31-a43a-4074bf7df29f', '2025-11-01 03:12:41.190826+00', '2025-11-01 03:12:41.190826+00', 'Zachary', 'Watsica', 'Deleniti eum laborum eos numquam non. Distinctio asperiores laudantium labore corporis est.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('c5afa4cc-8270-42ad-934d-a4087eb0dda9', '2025-11-01 03:12:41.190826+00', '2025-11-01 03:12:41.190826+00', 'Rickie', 'Rogahn', 'Ut voluptatum quo quod quibusdam quas. Velit blanditiis aut doloremque quam voluptas. Enim tenetur quibusdam est et ut. Corporis quos perspiciatis molestiae ut eum. Aspernatur id repudiandae quo commodi vel. Laborum neque ea voluptatum est quae. Qui cumque tenetur explicabo fuga repellendus.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('3f70f3e2-2a77-45d7-8f8b-d53a14ba6ff7', '2025-11-01 03:12:41.190826+00', '2025-11-01 03:12:41.190826+00', 'Alena', 'Eichmann', 'In aut et sit nihil qui.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('b46f301d-8fdc-421d-9b64-e39d7df44b15', '2025-11-01 03:12:41.190826+00', '2025-11-01 03:12:41.190826+00', 'Mabel', 'O''Reilly', 'Quos dolor et nulla dolor laboriosam. Itaque soluta illo qui ut laboriosam. Ipsum et dolores et reiciendis et. Quo quod occaecati ducimus molestiae autem. Placeat et adipisci odit maiores eos.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('5bdad376-22f1-4a8e-8139-fb6fb5b9b840', '2025-11-01 03:14:38.668989+00', '2025-11-01 03:14:38.668989+00', 'Ignatius', 'Bogan', 'Velit exercitationem recusandae dignissimos hic eum. Qui tempora tempore delectus quae iusto. Necessitatibus voluptates pariatur commodi aliquam a. Totam explicabo mollitia saepe nulla in. Laboriosam voluptatum vitae dolore cum ratione. Ut eum veritatis sed consequatur dolorem. Aut sequi itaque qui optio neque.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('d1a6d113-2c6e-43c2-8a18-3942d2e14176', '2025-11-01 03:14:38.668989+00', '2025-11-01 03:14:38.668989+00', 'Frank', 'Stiedemann', 'Architecto officia atque qui repellendus itaque. Voluptatem dolorum illo quidem quia laboriosam. Illo ad atque minus voluptas rerum. Sint rem quam dicta aspernatur maiores. Doloribus animi aliquid sit sapiente reiciendis. Quia debitis et omnis ratione voluptatum.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('c7728fbd-a4df-425d-8eb5-7f07841a5738', '2025-11-01 03:14:38.668989+00', '2025-11-01 03:14:38.668989+00', 'Carmel', 'Wisozk', 'Aut et placeat cum voluptate nemo. Consequuntur quos ex cum quis repudiandae. Aut dicta consequatur incidunt atque rerum. Explicabo temporibus rerum harum ipsum occaecati. Nesciunt autem et sint corrupti quidem. Sint illo consequatur aliquam ipsam iste.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('f1b4b278-596b-43c3-a02e-278884a0b9d6', '2025-11-01 03:14:38.668989+00', '2025-11-01 03:14:38.668989+00', 'Daisha', 'Abshire', 'Fugiat quidem iste dolores rem at. Numquam dolorum hic cumque molestiae iste. Laboriosam qui deserunt dolores nam at. Maiores occaecati et omnis incidunt et. Et ipsum quam assumenda dignissimos quia.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('f0fe8024-0d51-45f0-af8a-6bc7289f99e2', '2025-11-01 03:14:38.668989+00', '2025-11-01 03:14:38.668989+00', 'Sister', 'McGlynn', 'Asperiores quae porro qui facere est. Et est sint explicabo vel et. Sit et est impedit dolores dolorem. Ratione voluptatem velit sunt autem at. Quas dolorem nostrum sed esse est. Consequuntur molestiae voluptatibus esse ut laboriosam. Molestias dolorum sit quo ipsum excepturi. Quia voluptatem sit similique et doloremque. Tempore mollitia quidem quisquam repellat recusandae. Impedit in voluptatum odio magni eum.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('b4679309-69a2-408d-b5ce-b973105013d7', '2025-11-01 03:54:42.73476+00', '2025-11-01 03:54:42.73476+00', 'Christine', 'Waelchi', 'Omnis incidunt sapiente soluta at odio. Minus porro et et ipsum quod.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('2588cd72-0012-49c9-b6a6-d8e4e0644e83', '2025-11-01 03:54:42.73476+00', '2025-11-01 03:54:42.73476+00', 'Glenna', 'Labadie', 'Est non dolor et asperiores quos. Doloribus sint iusto doloremque fugit nobis. Sequi sunt totam natus est earum. Occaecati possimus sint eius commodi nisi. Et quo aut porro ipsam aut. Animi amet inventore aut similique assumenda. Qui perspiciatis minus inventore blanditiis quia.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('fe5e5818-fdfb-481c-aa71-00469de4764b', '2025-11-01 03:54:42.73476+00', '2025-11-01 03:54:42.73476+00', 'Hope', 'Okuneva', 'Ratione dolorem occaecati amet cum id. Eveniet nisi non quas quo odio. Necessitatibus fugiat et dolor consectetur enim.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('7dbb7a48-0d73-4e4e-8260-24c49871a347', '2025-11-01 03:54:42.73476+00', '2025-11-01 03:54:42.73476+00', 'Jane', 'Feeney', 'Velit est earum consequuntur nemo vero. Cupiditate dignissimos nobis dolorem aliquam consequatur.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('560f14b3-957b-4333-a9be-16ef4e4a8876', '2025-11-01 03:54:42.73476+00', '2025-11-01 03:54:42.73476+00', 'Rowena', 'Hamill', 'Aut ducimus ex exercitationem alias id. Quia perferendis reprehenderit ab reiciendis est. Quia fugit sint quibusdam repellat perspiciatis. Qui ea distinctio voluptatem sint inventore. Sit laboriosam quaerat dolor magnam repudiandae.');
INSERT INTO public.authors (id, created_at, updated_at, first_name, last_name, biography) VALUES ('0cabfefc-6729-452b-beba-456fa63b555d', '2025-11-01 05:06:18.530767+00', '2025-11-01 05:07:07.955696+00', 'Zukron', 'Alviandy Rahmadhan', 'Hello World');


--
-- TOC entry 3441 (class 0 OID 158201)
-- Dependencies: 220
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('c767efd6-70cd-4eff-9e51-cab58c8d851e', '2025-11-01 03:54:42.824689+00', '2025-11-01 03:54:42.824689+00', 'Miss Leanna Howe', 'Quos molestiae eos omnis aliquid numquam. Eaque quasi deserunt laborum dolore voluptatum.', '3659804217', '2025-11-01 03:54:42.801987+00', '6e1cdcf7-6711-464f-a7c1-e3e02f885d76');
INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('5932aa91-7852-4602-9397-8de4c2077141', '2025-11-01 03:54:42.824689+00', '2025-11-01 03:54:42.824689+00', 'Lady Zella Cremin', 'Laboriosam voluptatem maxime veniam mollitia qui. Explicabo enim pariatur quos illum non. Quas consequatur accusamus quia rerum commodi. Est doloribus sit exercitationem excepturi molestias. Ut culpa neque ea doloribus ex. Aperiam voluptatibus dolor occaecati corrupti natus. Nisi perspiciatis cum et voluptas aliquam. Minus temporibus adipisci porro exercitationem rem.', '9265183740', '2025-11-01 03:54:42.802067+00', '746e25d7-697e-42c5-bce4-5f9d6e984bb1');
INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('c5002f8d-4d05-4550-b640-4351382da269', '2025-11-01 03:54:42.824689+00', '2025-11-01 03:54:42.824689+00', 'Prof. Cathrine Bogisich', 'Quibusdam aut neque ea et quia. Repudiandae voluptate optio necessitatibus libero non. Ut sunt consequuntur voluptatem perspiciatis dolorum. Velit voluptates aspernatur nam molestiae vitae.', '1364527809', '2025-11-01 03:54:42.802109+00', 'd165ed70-b2c5-4b48-92ec-5a5202f70e06');
INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('d82af8d3-b24b-4da0-a3a4-a712fc4b8823', '2025-11-01 03:54:42.824689+00', '2025-11-01 03:54:42.824689+00', 'Dr. Tressa Ankunding', 'Et reprehenderit ut temporibus voluptatem quo. Corporis necessitatibus aut totam magnam adipisci. Repellendus maiores ea velit sit culpa. Qui dolore voluptatem itaque nostrum molestias. Consequatur non sint rerum a aut. Voluptatem eum est assumenda vel voluptatem. Voluptas vero eligendi animi et enim. Aut eos laborum reiciendis reprehenderit expedita. Quasi accusantium enim unde officia ut. Libero impedit praesentium sunt dolor esse.', '5891476032', '2025-11-01 03:54:42.802169+00', 'fc67dc73-902b-445c-a828-c4be32367172');
INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('aaf63bb8-4d0f-4550-bd3d-d4fd40d3506d', '2025-11-01 03:54:42.824689+00', '2025-11-01 03:54:42.824689+00', 'Prof. Marguerite Medhurst', 'Qui fugiat et quidem nulla omnis. Quo nihil neque optio dolorem id.', '2398156047', '2025-11-01 03:54:42.802193+00', '3153eb3f-4ad9-4dd2-9efb-18f0b45d226b');
INSERT INTO public.books (id, created_at, updated_at, title, description, isbn, published_at, publisher_id) VALUES ('1cd15f16-c533-4cc3-ae1f-58dc26e678c8', '2025-11-01 08:42:44.464354+00', '2025-11-01 08:42:44.464354+00', 'Naga Bonar', 'Naga Bonar Bertapa di Goa Hantu', '12345', '2025-11-01 06:09:17+00', '746e25d7-697e-42c5-bce4-5f9d6e984bb1');


--
-- TOC entry 3440 (class 0 OID 158193)
-- Dependencies: 219
-- Data for Name: publishers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('6e1cdcf7-6711-464f-a7c1-e3e02f885d76', '2025-11-01 03:14:38.725424+00', '2025-11-01 03:14:38.725424+00', 'Mr. Norbert Swift', '140 South Hill Avenue', '269-871-0543', 'http://yvtzvtq.com/pwpVPKJ');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('fc67dc73-902b-445c-a828-c4be32367172', '2025-11-01 03:14:38.725424+00', '2025-11-01 03:14:38.725424+00', 'Lord Clyde Von', '736 Sicard Street Southeast', '573-104-2619', 'https://www.axywbch.com/TwFDvAj');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('fff80dfd-3fea-4f4f-9bb8-89dfb16381c2', '2025-11-01 03:14:38.725424+00', '2025-11-01 03:14:38.725424+00', 'Prof. Roger Goyette', '14 Huntington Street', '248-710-9536', 'http://www.wujdhws.net/');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('42d78aff-7743-4145-beee-c77aa2804225', '2025-11-01 03:14:38.725424+00', '2025-11-01 03:14:38.725424+00', 'King Angelo Corwin', '8467 Chase Drive', '174-685-3910', 'https://zkboywx.edu/JlVcjsG.html');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('112ebdae-bab1-4804-afd7-e8f7042ae172', '2025-11-01 03:54:42.758391+00', '2025-11-01 03:54:42.758391+00', 'Lady Raegan Muller', '6398 Jellison Way', '257-911-0348', 'http://otidiwi.biz/hYlEyJY.css');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('d165ed70-b2c5-4b48-92ec-5a5202f70e06', '2025-11-01 03:54:42.758391+00', '2025-11-01 03:54:42.758391+00', 'Ms. Nyah Rodriguez', '5843 West McLellan Road', '952-876-4131', 'http://vbglysy.top/UOFjhCQ.jpg');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('76dd3131-3f20-4207-9968-f951152279de', '2025-11-01 03:54:42.758391+00', '2025-11-01 03:54:42.758391+00', 'Miss Ashlynn Smitham', '7002 Secrest Court', '103-274-9815', 'https://tkksljq.biz/iUOGuhn.js');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('3153eb3f-4ad9-4dd2-9efb-18f0b45d226b', '2025-11-01 03:54:42.758391+00', '2025-11-01 03:54:42.758391+00', 'Queen Samantha Crona', '119 Oakland Street', '597-102-1346', 'http://xbsdoly.net/XgYkKYX.webp');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('54d072a4-3079-42d9-95ae-230f705c69b6', '2025-11-01 05:18:10.09823+00', '2025-11-01 05:18:10.09823+00', 'Gramedia', 'Yogyakarta', '0895123', 'https://gramedia.com');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('746e25d7-697e-42c5-bce4-5f9d6e984bb1', '2025-11-01 03:14:38.725424+00', '2025-11-01 05:18:33.47433+00', 'Tiga Serangkai', '19590 East Batavia Drive', '471-062-5891', 'http://dycokwl.org/UMOfPMi.html');
INSERT INTO public.publishers (id, created_at, updated_at, name, address, phone, website) VALUES ('94c93e27-2778-4515-994b-33722694bfeb', '2025-11-01 06:28:47.673497+00', '2025-11-01 06:28:47.673497+00', 'Gramedia', 'Yogyakarta', '0895123', 'https://gramedia.com');


--
-- TOC entry 3443 (class 0 OID 163043)
-- Dependencies: 222
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, created_at, updated_at, username, password) VALUES ('bc369fd1-35e3-4c30-a0d8-b729096b073a', '2025-11-01 08:21:52.600171+00', '2025-11-01 08:21:52.600171+00', 'heyaa', '$2a$10$ebNk50d34PqO1wvlcGu7Iebb30q4Z/g3osjWPMnqM4XztjKW8BXfG');


--
-- TOC entry 3288 (class 2606 OID 158221)
-- Name: author_books author_books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.author_books
    ADD CONSTRAINT author_books_pkey PRIMARY KEY (book_id, author_id);


--
-- TOC entry 3281 (class 2606 OID 158192)
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- TOC entry 3285 (class 2606 OID 158208)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 3283 (class 2606 OID 158200)
-- Name: publishers publishers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.publishers
    ADD CONSTRAINT publishers_pkey PRIMARY KEY (id);


--
-- TOC entry 3290 (class 2606 OID 163050)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3286 (class 1259 OID 158214)
-- Name: idx_books_publisher_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_books_publisher_id ON public.books USING btree (publisher_id);


--
-- TOC entry 3292 (class 2606 OID 158227)
-- Name: author_books fk_author_books_author; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.author_books
    ADD CONSTRAINT fk_author_books_author FOREIGN KEY (author_id) REFERENCES public.authors(id) ON DELETE CASCADE;


--
-- TOC entry 3293 (class 2606 OID 158222)
-- Name: author_books fk_author_books_book; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.author_books
    ADD CONSTRAINT fk_author_books_book FOREIGN KEY (book_id) REFERENCES public.books(id) ON DELETE CASCADE;


--
-- TOC entry 3291 (class 2606 OID 158209)
-- Name: books fk_publishers_book; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT fk_publishers_book FOREIGN KEY (publisher_id) REFERENCES public.publishers(id) ON UPDATE CASCADE ON DELETE RESTRICT;


-- Completed on 2025-11-01 08:52:58 UTC

--
-- PostgreSQL database dump complete
--

